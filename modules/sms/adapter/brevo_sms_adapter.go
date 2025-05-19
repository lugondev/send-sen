package adapter

import (
	"context"
	"fmt"

	"github.com/lugondev/send-sen/modules/sms/port"
	"github.com/lugondev/send-sen/pkg/logger"
	// TODO: Import Brevo SDK if available and needed for SMS
)

// BrevoSMSAdapter implements smsPort.SMSAdapter and ports.HealthChecker for Brevo SMS.
type BrevoSMSAdapter struct {
	apiKey      string
	senderName  string // Optional: Sender name for SMS
	logger      logger.Logger
	serviceName string
	// Brevo client if using SDK
}

// BrevoSMSConfig holds configuration for the Brevo SMS adapter.
// Ensure mapstructure tags match the config file structure.
type BrevoSMSConfig struct {
	APIKey     string `mapstructure:"api_key"`
	SenderName string `mapstructure:"sender_name"` // Optional alphanumeric sender ID
}

// NewBrevoSMSAdapter creates a new Brevo SMS adapter.
func NewBrevoSMSAdapter(config BrevoSMSConfig, logger logger.Logger) (port.SMSAdapter, error) {
	if config.APIKey == "" {
		return nil, fmt.Errorf("brevo API key is required for SMS adapter")
	}
	namedLogger := logger.WithFields(map[string]any{
		"service": "brevo_sms",
	})
	adapter := &BrevoSMSAdapter{
		apiKey:      config.APIKey,
		senderName:  config.SenderName,
		logger:      namedLogger,
		serviceName: "brevo_sms",
	}
	ctx := context.Background()
	namedLogger.Info(ctx, "Brevo SMS adapter initialized", map[string]any{
		"service_name": adapter.serviceName,
		"sender_name":  adapter.senderName,
	})
	return adapter, nil
}

// SendSMS sends an SMS using the Brevo API. Matches the smsPort.SMSAdapter interface.
func (a *BrevoSMSAdapter) SendSMS(sms port.SMS) error {
	// Use the sender name configured for the adapter if sms.From is empty
	sender := a.senderName
	if sms.From != "" {
		sender = sms.From // Allow overriding sender per message if provided
	}

	a.logger.Info(context.Background(), "Attempting to send SMS via Brevo", map[string]any{
		"to":             sms.To,
		"message_length": len(sms.Message),
		"sender":         sender,
	})

	// --- Placeholder for Brevo SMS API Call ---
	// TODO: Implement the actual API call to Brevo to send the SMS.
	// This might involve using the Brevo SDK's transactional SMS endpoint
	// or making a direct HTTP request.
	/*
		Example using hypothetical SDK:
		smsData := brevoSDK.SendTransacSms{
			Sender:    &sender, // Use the determined sender
			Recipient: &sms.To,
			Content:   &sms.Message,
			// Type: "transactional", // or "marketing"
		}
		// Note: Context is not passed here, but SDK methods might take it. Adjust if needed.
		_, response, err := a.client.TransactionalSMSApi.SendTransacSms(context.Background(), smsData) // Pass a background context for now
		if err != nil {
			// Handle error
			return fmt.Errorf("brevo SMS API error: %w", err)
		}
		if response.StatusCode >= 300 {
			// Handle non-success status
			return fmt.Errorf("brevo SMS API returned status code %d", response.StatusCode)
		}
	*/

	a.logger.Warn(context.Background(), "Brevo SendSMS function is not fully implemented yet.")
	fmt.Printf("--- MOCK Brevo SMS Send ---\nTo: %s\nSender: %s\nMessage: %s\n-------------------------\n",
		sms.To, sender, sms.Message)
	// --- End Placeholder ---

	return nil // Return nil for now
}

// ServiceName implements the ports.HealthChecker interface.
func (a *BrevoSMSAdapter) ServiceName() string {
	return a.serviceName
}
