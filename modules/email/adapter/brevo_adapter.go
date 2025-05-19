package adapter

import (
	"context"
	"fmt"

	brevo "github.com/getbrevo/brevo-go/lib"
	"github.com/lugondev/send-sen/config"
	"github.com/lugondev/send-sen/modules/email/port"
	"github.com/lugondev/send-sen/pkg/logger"
	"github.com/samber/lo"
)

// BrevoAdapter implements the port.EmailAdapter interface for sending emails via Brevo (formerly SendinBlue).
type BrevoAdapter struct {
	apiKey string
	logger logger.Logger
	client *brevo.APIClient
	cfg    config.BrevoConfig
}

// NewBrevoAdapter creates a new instance of BrevoAdapter.
func NewBrevoAdapter(cfg config.BrevoConfig, logger logger.Logger) (port.EmailAdapter, error) {
	if cfg.APIKey == "" {
		return nil, fmt.Errorf("brevo API key is required")
	}
	namedLogger := logger.WithFields(map[string]any{
		"service": "brevo_email",
	})

	brevoCfg := brevo.NewConfiguration()
	// Configure API key authorization
	brevoCfg.AddDefaultHeader("api-key", cfg.APIKey)
	// Create new API client
	apiClient := brevo.NewAPIClient(brevoCfg)

	adapter := &BrevoAdapter{
		apiKey: cfg.APIKey,
		logger: namedLogger,
		client: apiClient,
		cfg:    cfg,
	}
	namedLogger.Info(context.Background(), "Brevo email adapter initialized")
	return adapter, nil
}

// SendEmail sends an email using the Brevo API.
func (a *BrevoAdapter) SendEmail(ctx context.Context, email port.Email) error {
	a.logger.Info(ctx, "Attempting to send email via Brevo", map[string]any{
		"subject": email.Subject,
		"to":      email.To,
		"cc":      email.Cc,
		"bcc":     email.Bcc,
	})

	// Create a new email
	sendSmtpEmail := brevo.SendSmtpEmail{
		// Set sender information
		Sender: &brevo.SendSmtpEmailSender{
			Name:  a.cfg.SenderName,
			Email: a.cfg.SenderEmail,
		},
		// Set recipient information
		To: lo.Map(email.To, func(to string, _ int) brevo.SendSmtpEmailTo {
			return brevo.SendSmtpEmailTo{
				Email: to,
				Name:  to,
			}
		}),
		// Set email subject
		Subject: email.Subject,
		// Set email content (HTML)
		HtmlContent: email.Html,
		// Optional: Set plain text content for email clients that don't support HTML
		TextContent: email.Body,
	}

	// Send the email
	result, response, err := a.client.TransactionalEmailsApi.SendTransacEmail(ctx, sendSmtpEmail)
	// Check if there was an error
	if err != nil {
		fmt.Printf("Error sending email: %s\n", err.Error())
		return err
	}

	// Print the result
	fmt.Printf("Email sent successfully! Message ID: %s\n", result.MessageId)
	fmt.Printf("API Response Status: %s\n", response.Status)

	return nil
}
