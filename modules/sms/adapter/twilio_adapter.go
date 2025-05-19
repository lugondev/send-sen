package adapter

import (
	"context"
	"fmt"
	"net/url"

	"github.com/lugondev/send-sen/modules/sms/port"
	"github.com/lugondev/send-sen/pkg/logger"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

// TwilioAdapter implements the port.SMSAdapter and ports.HealthChecker interfaces using the Twilio API.
type TwilioAdapter struct {
	client       *twilio.RestClient
	accountSid   string
	authToken    string
	messagingSid string // Optional: Specific Messaging Service SID
	defaultFrom  string // Optional: Default 'From' number if Messaging SID not used
	logger       logger.Logger
	serviceName  string
}

// TwilioConfig holds the configuration needed for the Twilio adapter.
type TwilioConfig struct {
	AccountSid   string `mapstructure:"account_sid"`   // Required
	AuthToken    string `mapstructure:"auth_token"`    // Required
	MessagingSid string `mapstructure:"messaging_sid"` // Optional: Recommended over 'From' number
	FromNumber   string `mapstructure:"from_number"`   // Optional: Required if MessagingSid is not set
}

// NewTwilioAdapter creates a new instance of TwilioAdapter.
// Returns both SMS adapter and health checker interfaces.
func NewTwilioAdapter(config TwilioConfig, logger logger.Logger) (port.SMSAdapter, error) {
	if config.AccountSid == "" || config.AuthToken == "" {
		return nil, fmt.Errorf("twilio Account SID and Auth Token are required")
	}
	if config.MessagingSid == "" && config.FromNumber == "" {
		return nil, fmt.Errorf("either Twilio Messaging Service SID or From Number is required")
	}
	namedLogger := logger.WithFields(map[string]any{
		"service":       "twilio_sms",
		"account_sid":   config.AccountSid,
		"messaging_sid": config.MessagingSid,
		"from_number":   config.FromNumber,
	})
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		AccountSid: config.AccountSid,
		Password:   config.AuthToken,
	})
	adapter := &TwilioAdapter{
		client:       client,
		accountSid:   config.AccountSid,
		authToken:    config.AuthToken,
		messagingSid: config.MessagingSid,
		defaultFrom:  config.FromNumber,
		logger:       namedLogger,
		serviceName:  "twilio_sms",
	}
	ctx := context.Background()
	namedLogger.Info(ctx, "Twilio SMS adapter initialized")
	return adapter, nil
}

// SendSMS sends an SMS using the Twilio Messages API.
func (a *TwilioAdapter) SendSMS(sms port.SMS) error {
	params := &twilioApi.CreateMessageParams{}
	params.SetTo(sms.To)
	params.SetBody(sms.Message)

	// Determine the 'From' - prioritize Messaging Service SID
	if a.messagingSid != "" {
		params.SetMessagingServiceSid(a.messagingSid)
		a.logger.Debug(context.Background(), "Sending SMS via Twilio Messaging Service", map[string]any{"messaging_sid": a.messagingSid})
	} else {
		// Use specific 'From' from SMS struct if provided, otherwise use default 'From' number
		fromNumber := sms.From
		if fromNumber == "" {
			fromNumber = a.defaultFrom
		}
		if fromNumber == "" {
			// This check should be redundant due to NewTwilioAdapter validation, but safety first.
			return fmt.Errorf("cannot send Twilio SMS: missing MessagingServiceSid and From number")
		}
		params.SetFrom(fromNumber)
		a.logger.Debug(context.Background(), "Sending SMS via Twilio From Number", map[string]any{"from_number": fromNumber})
	}

	// Add other optional parameters if needed from sms struct
	// e.g., params.SetStatusCallback("http://example.com/status")

	a.logger.Info(context.Background(), "Attempting to send SMS via Twilio", map[string]any{
		"to":             sms.To,
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
		a.logger.Info(context.Background(), "SMS sent successfully via Twilio", map[string]any{
			"to":          sms.To,
			"message_sid": *resp.Sid,
			"status":      *resp.Status,
		})
	} else {
		// Should not happen if err is nil, but log just in case
		a.logger.Warn(context.Background(), "Twilio API call returned nil error but also nil response/SID")
	}

	return nil
}

// ServiceName implements the ports.HealthChecker interface.
func (a *TwilioAdapter) ServiceName() string {
	return a.serviceName
}
