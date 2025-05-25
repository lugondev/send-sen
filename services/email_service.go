package services

import (
	"bytes"
	"context"
	"fmt"
	"html/template"

	adapter "github.com/lugondev/send-sen/adapters/email"
	"github.com/lugondev/send-sen/dto"
	"github.com/lugondev/send-sen/ports"

	logger "github.com/lugondev/go-log"
	"github.com/lugondev/send-sen/config"
)

// emailService implements the Service interface.
type emailService struct {
	adapter ports.EmailAdapter
	logger  logger.Logger
	name    config.EmailProvider
}

// NewEmailService creates a new instance of Service.
func NewEmailService(cfg config.Config, logger logger.Logger) (ports.EmailService, error) {
	ctx := context.Background()
	logger.Debug(ctx, "Registered email adapter", map[string]any{
		"adapter": cfg.Adapter.Email,
	})

	var emailAdapter ports.EmailAdapter
	if cfg.Adapter.Email == config.EmailBrevo {
		brevoAdapter, err := adapter.NewBrevoAdapter(cfg.Brevo, logger)
		if err != nil {
			logger.Error(ctx, "Failed to create Brevo adapter", map[string]any{
				"error": err,
			})
		} else {
			emailAdapter = brevoAdapter
			logger.Info(ctx, "Using Brevo adapter for email sending")
		}
	} else if cfg.Adapter.Email == config.EmailSendGrid {
		sendgridAdapter, err := adapter.NewSendGridAdapter(cfg.SendGrid, logger)
		if err != nil {
			logger.Error(ctx, "Failed to create SendGrid adapter", map[string]any{
				"error": err,
			})
		} else {
			emailAdapter = sendgridAdapter
			logger.Info(ctx, "Using SendGrid adapter for email sending")
		}
	}
	if emailAdapter == nil {
		emailAdapter = adapter.NewMockEmailAdapter(logger)
		logger.Info(ctx, "Using MockEmail adapter for email sending")
	}

	return &emailService{
		adapter: emailAdapter,
		logger: logger.WithFields(map[string]any{
			"service": "email_service_" + cfg.Adapter.Email,
		}),
		name: cfg.Adapter.Email,
	}, nil
}

// SendEmail delegates the email sending task to the configured adapter.
func (s *emailService) SendEmail(ctx context.Context, message dto.Email) error {
	if len(message.To) == 0 {
		return fmt.Errorf("message must have at least one recipient")
	}
	if message.Subject == "" {
		return fmt.Errorf("message subject cannot be empty")
	}
	if message.Body == "" {
		return fmt.Errorf("message body cannot be empty")
	}

	// Delegate to the adapter
	err := s.adapter.SendEmail(ctx, message)
	if err != nil {
		// Log the error maybe?
		return fmt.Errorf("failed to send message via adapter: %w", err)
	}
	return nil
}

// ---------- private helpers ----------

// renderHTML compile template & return HTML.
func (s *emailService) renderHTML(tpl string, data any) (string, error) {
	t, err := template.New("").Parse(tpl)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (s *emailService) ServiceName() string {
	return string(s.name)
}

// SendPasswordReset sends a password reset email with a reset link.
func (s *emailService) SendPasswordReset(ctx context.Context, to string, link string) error {
	// Render the HTML template
	html, err := s.renderHTML(tplReset, map[string]string{
		"Link": link,
	})
	if err != nil {
		s.logger.Error(ctx, "Failed to render password reset template", map[string]any{
			"error": err,
		})
		return fmt.Errorf("failed to render password reset template: %w", err)
	}

	// Create the email message
	message := dto.Email{
		To:      []string{to},
		Subject: "Password Reset Request",
		Html:    html,
		Body:    "You have requested to reset your password. Click the link to continue: " + link,
	}

	// Send the email
	return s.SendEmail(ctx, message)
}

// SendVerificationCode sends an email with a verification code.
func (s *emailService) SendVerificationCode(ctx context.Context, to string, code string) error {
	// Render the HTML template
	html, err := s.renderHTML(tplVerificationCode, map[string]string{
		"Code": code,
	})
	if err != nil {
		s.logger.Error(ctx, "Failed to render verification code template", map[string]any{
			"error": err,
		})
		return fmt.Errorf("failed to render verification code template: %w", err)
	}

	// Create the email message
	message := dto.Email{
		To:      []string{to},
		Subject: "Your Verification Code",
		Html:    html,
		Body:    "Your verification code is: " + code + ". This code will expire in 10 minutes.",
	}

	// Send the email
	return s.SendEmail(ctx, message)
}

// SendWelcome sends a welcome email to a new user.
func (s *emailService) SendWelcome(ctx context.Context, to string, name string) error {
	// Render the HTML template
	html, err := s.renderHTML(tplWelcome, map[string]string{
		"Name": name,
	})
	if err != nil {
		s.logger.Error(ctx, "Failed to render welcome template", map[string]any{
			"error": err,
		})
		return fmt.Errorf("failed to render welcome template: %w", err)
	}

	// Create the email message
	message := dto.Email{
		To:      []string{to},
		Subject: "Welcome to MyService!",
		Html:    html,
		Body:    "Hello " + name + ", Welcome to MyService! Explore our amazing features right now.",
	}

	// Send the email
	return s.SendEmail(ctx, message)
}

// SendWarningLogin sends a warning email about a new login from an unfamiliar location.
func (s *emailService) SendWarningLogin(ctx context.Context, to string, location string, time string) error {
	// Render the HTML template
	html, err := s.renderHTML(tplWarningLogin, map[string]string{
		"Location": location,
		"Time":     time,
	})
	if err != nil {
		s.logger.Error(ctx, "Failed to render login warning template", map[string]any{
			"error": err,
		})
		return fmt.Errorf("failed to render login warning template: %w", err)
	}

	// Create the email message
	message := dto.Email{
		To:      []string{to},
		Subject: "Security Alert: New Login Detected",
		Html:    html,
		Body:    "We detected a new login to your account from " + location + " at " + time + ". If this wasn't you, please secure your account immediately.",
	}

	// Send the email
	return s.SendEmail(ctx, message)
}
