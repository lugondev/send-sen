package sms

import (
	"context"
	"fmt"

	brevo "github.com/getbrevo/brevo-go/lib"
	"github.com/lugondev/go-log"
	"github.com/lugondev/send-sen/config"
)

// BrevoAdapter implements the port.SmsAdapter interface for sending SMS via Brevo (formerly SendinBlue).
type BrevoAdapter struct {
	apiKey string
	logger logger.Logger
	client *brevo.APIClient
	cfg    config.BrevoConfig
}

// NewBrevoAdapter creates a new instance of BrevoAdapter.
func NewBrevoAdapter(cfg config.BrevoConfig, logger logger.Logger) (*BrevoAdapter, error) {
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

// Send sends an SMS using the Brevo API.
func (a *BrevoAdapter) Send(ctx context.Context, sms SMS) error {
	a.logger.Info(ctx, "Attempting to send SMS via Brevo", map[string]any{
		"to":      sms.To,
		"from":    a.cfg.SMSSender,
		"message": sms.Message,
	})

	sendTransacSms := brevo.SendTransacSms{
		// Sender - company or brand name (alphanumeric, max 11 chars) or phone number
		Sender: a.cfg.SMSSender,
		// Recipient's phone number with country code (e.g., "33612345678" for France)
		Recipient: sms.To,
		// SMS content
		Content: sms.Message,
	}

	// Send the SMS
	if _, _, err := a.client.TransactionalSMSApi.SendTransacSms(ctx, sendTransacSms); err != nil {
		fmt.Printf("Error sending SMS: %s\n", err.Error())
		return err
	}

	return nil
}
