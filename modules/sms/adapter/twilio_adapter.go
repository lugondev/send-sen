package adapter

import (
	"context"
	"fmt"
	"net/url"

	"github.com/lugondev/send-sen/config"
	"github.com/lugondev/send-sen/modules/sms/port"
	"github.com/lugondev/send-sen/pkg/logger"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

// TwilioAdapter implements the port.SMSAdapter and ports.HealthChecker interfaces using the Twilio API.
type TwilioAdapter struct {
	client      *twilio.RestClient
	cfg         config.TwilioConfig
	logger      logger.Logger
	serviceName string
}

// NewTwilioAdapter creates a new instance of TwilioAdapter.
// Returns both SMS adapter and health checker interfaces.
func NewTwilioAdapter(cfg config.TwilioConfig, logger logger.Logger) (port.SMSAdapter, error) {
	if cfg.AccountSid == "" || cfg.AuthToken == "" {
		return nil, fmt.Errorf("twilio Account SID and Auth Token are required")
	}
	if cfg.MessagingSid == "" && cfg.FromNumber == "" {
		return nil, fmt.Errorf("either Twilio Messaging Service SID or From Number is required")
	}
	namedLogger := logger.WithFields(map[string]any{
		"service":       "twilio_sms",
		"account_sid":   cfg.AccountSid,
		"messaging_sid": cfg.MessagingSid,
		"from_number":   cfg.FromNumber,
	})
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cfg.AccountSid,
		Password: cfg.AuthToken,
	})
	adapter := &TwilioAdapter{
		client:      client,
		cfg:         cfg,
		logger:      namedLogger,
		serviceName: "twilio_sms",
	}
	ctx := context.Background()
	namedLogger.Info(ctx, "Twilio SMS adapter initialized")
	return adapter, nil
}

// SendSMS sends an SMS using the Twilio Messages API.
func (a *TwilioAdapter) SendSMS(ctx context.Context, sms port.SMS) error {
	params := &twilioApi.CreateMessageParams{
		To:   &sms.To,
		From: &a.cfg.FromNumber,
		Body: &sms.Message,
	}

	a.logger.Info(ctx, "Attempting to send SMS via Twilio", map[string]any{
		"to":             sms.To,
		"from":           a.cfg.FromNumber,
		"message_length": len(sms.Message),
	})

	// Send the message
	resp, err := a.client.Api.CreateMessage(params)
	if err != nil {
		// Handle specific Twilio errors if possible (e.g., using url.Error)
		if urlErr, ok := err.(*url.Error); ok {
			a.logger.Error(context.Background(), "Twilio API request network error", map[string]any{"error": urlErr})
			return fmt.Errorf("twilio network error: %w", err)
		}
		// Generic error handling
		a.logger.Error(context.Background(), "Failed to send SMS via Twilio API", map[string]any{"error": err})
		return fmt.Errorf("twilio API error: %w", err)
	}

	// Check response details (resp will be nil on error)
	// Twilio library handles standard success/failure via the error return.
	// We can log the SID of the created message for tracking.
	if resp != nil && resp.Sid != nil {
		a.logger.Info(ctx, "SMS sent successfully via Twilio", map[string]any{
			"to":          sms.To,
			"message_sid": *resp.Sid,
			"status":      *resp.Status,
		})
	} else {
		// Should not happen if err is nil, but log just in case
		a.logger.Warn(ctx, "Twilio API call returned nil error but also nil response/SID")
	}

	return nil
}

// ServiceName implements the ports.HealthChecker interface.
func (a *TwilioAdapter) ServiceName() string {
	return a.serviceName
}
