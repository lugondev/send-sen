package service

import (
	"context"
	"fmt"

	"github.com/lugondev/send-sen/config"
	"github.com/lugondev/send-sen/modules/email/adapter"
	"github.com/lugondev/send-sen/modules/email/port"
	"github.com/lugondev/send-sen/pkg/logger"
)

// emailService implements the EmailService interface.
type emailService struct {
	adapter port.EmailAdapter
	logger  logger.Logger
	name    config.EmailProvider
}

// NewEmailService creates a new instance of EmailService.
func NewEmailService(cfg config.Config, logger logger.Logger) port.EmailService {
	namedLogger := logger.WithFields(map[string]any{
		"service": "email_service_" + cfg.Adapter.Email,
	})
	ctx := context.Background()
	namedLogger.Debug(ctx, "Registered email adapter", map[string]any{
		"adapter": cfg.Adapter.Email,
	})
	var emailAdapter port.EmailAdapter
	if cfg.Adapter.Email == config.EmailBrevo {
		brevoAdapter, err := adapter.NewBrevoAdapter(cfg.Brevo, namedLogger)
		if err != nil {
			namedLogger.Error(ctx, "Failed to create Brevo adapter", map[string]any{
				"error": err,
			})
		} else {
			emailAdapter = brevoAdapter
			namedLogger.Info(ctx, "Using Brevo adapter for email sending")
		}
	} else if cfg.Adapter.Email == config.EmailSendGrid {
		sendgridAdapter, err := adapter.NewSendGridAdapter(cfg.SendGrid, namedLogger)
		if err != nil {
			namedLogger.Error(ctx, "Failed to create SendGrid adapter", map[string]any{
				"error": err,
			})
		} else {
			emailAdapter = sendgridAdapter
			namedLogger.Info(ctx, "Using SendGrid adapter for email sending")
		}
	}
	if emailAdapter == nil {
		emailAdapter = adapter.NewMockEmailAdapter(namedLogger)
		namedLogger.Info(ctx, "Using MockEmail adapter for email sending")
	}

	return &emailService{
		adapter: emailAdapter,
		logger:  logger,
	}
}

// SendEmail delegates the email sending task to the configured adapter.
func (s *emailService) SendEmail(ctx context.Context, email port.Email) error {
	if len(email.To) == 0 {
		return fmt.Errorf("email must have at least one recipient")
	}
	if email.Subject == "" {
		return fmt.Errorf("email subject cannot be empty")
	}
	if email.Body == "" {
		return fmt.Errorf("email body cannot be empty")
	}

	// Delegate to the adapter
	err := s.adapter.SendEmail(ctx, email)
	if err != nil {
		// Log the error maybe?
		return fmt.Errorf("failed to send email via adapter: %w", err)
	}
	return nil
}

func (s *emailService) ServiceName() string {
	return string(s.name)
}
