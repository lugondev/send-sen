package adapter

import (
	"context"
	"fmt"

	"github.com/lugondev/send-sen/config"
	"github.com/lugondev/send-sen/modules/email/port"
	"github.com/lugondev/send-sen/pkg/logger"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendGridAdapter implements the port.EmailAdapter and ports.HealthChecker interfaces using the SendGrid API.
type SendGridAdapter struct {
	apiKey      string
	client      *sendgrid.Client
	from        *mail.Email
	logger      logger.Logger
	serviceName string
}

// NewSendGridAdapter creates a new instance of SendGridAdapter.
// Returns both email adapter and health checker interfaces.
func NewSendGridAdapter(cfg config.SendGridConfig, logger logger.Logger) (port.EmailAdapter, error) {
	if cfg.APIKey == "" {
		return nil, fmt.Errorf("SendGrid API key is required")
	}
	if cfg.FromEmail == "" {
		return nil, fmt.Errorf("SendGrid From email address is required")
	}
	namedLogger := logger.WithFields(map[string]any{
		"service": "sendgrid_email",
	})
	client := sendgrid.NewSendClient(cfg.APIKey)
	from := mail.NewEmail(cfg.FromName, cfg.FromEmail)
	adapter := &SendGridAdapter{
		apiKey:      cfg.APIKey,
		client:      client,
		from:        from,
		logger:      namedLogger,
		serviceName: "sendgrid_email",
	}
	ctx := context.Background()
	namedLogger.Info(ctx, "SendGrid email adapter initialized", map[string]any{
		"service_name": adapter.serviceName,
		"from_email":   cfg.FromEmail,
		"from_name":    cfg.FromName,
	})
	return adapter, nil
}

// SendEmail sends an email using the SendGrid API.
func (a *SendGridAdapter) SendEmail(ctx context.Context, email port.Email) error {
	a.logger.Info(ctx, "Attempting to send email via SendGrid", map[string]any{
		"subject": email.Subject,
		"to":      email.To,
		"cc":      email.Cc,
		"bcc":     email.Bcc,
		"from":    a.from.Address,
	})

	// Create SendGrid message structure
	message := mail.NewV3Mail()
	message.SetFrom(a.from)
	message.Subject = email.Subject

	// Create personalization block for To, Cc, Bcc
	p := mail.NewPersonalization()

	// Add To recipients
	toEmails := make([]*mail.Email, 0, len(email.To))
	for _, recipient := range email.To {
		toEmails = append(toEmails, mail.NewEmail("", recipient)) // Name can be empty
	}
	p.AddTos(toEmails...)

	// Add Cc recipients
	if len(email.Cc) > 0 {
		ccEmails := make([]*mail.Email, 0, len(email.Cc))
		for _, recipient := range email.Cc {
			ccEmails = append(ccEmails, mail.NewEmail("", recipient))
		}
		p.AddCCs(ccEmails...)
	}

	// Add Bcc recipients
	if len(email.Bcc) > 0 {
		bccEmails := make([]*mail.Email, 0, len(email.Bcc))
		for _, recipient := range email.Bcc {
			bccEmails = append(bccEmails, mail.NewEmail("", recipient))
		}
		p.AddBCCs(bccEmails...)
	}

	message.AddPersonalizations(p)

	htmlContent := email.Body
	message.AddContent(mail.NewContent("text/html", htmlContent))

	// Send the email
	_, err := a.client.Send(message)
	if err != nil {
		a.logger.Error(context.Background(), "Failed to send email via SendGrid API", map[string]any{
			"error": err,
		})
		return fmt.Errorf("sendGrid API error: %w", err)
	}
	a.logger.Info(ctx, "Email sent successfully via SendGrid", map[string]any{
		"subject": email.Subject,
		"to":      email.To,
		"content": htmlContent,
	})

	return nil
}

// ServiceName implements the ports.HealthChecker interface.
func (a *SendGridAdapter) ServiceName() string {
	return a.serviceName
}
