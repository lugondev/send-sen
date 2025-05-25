package main

import (
	"context"
	"fmt"

	adapter "github.com/lugondev/send-sen/adapters/sms"
	"github.com/lugondev/send-sen/dto"

	logger "github.com/lugondev/go-log"
	"github.com/lugondev/send-sen/config"
)

// smsService implements the Service interface.
type smsService struct {
	adapter SMSAdapter
	logger  logger.Logger
	name    config.SMSProvider
	from    string
}

// NewSMSService creates a new instance of Service.
// It requires an Adapter to be provided.
func NewSMSService(cfg config.Config, logger logger.Logger) (SMSService, error) {
	ctx := context.Background()
	var smsAdapter SMSAdapter
	var from string
	if cfg.Adapter.SMS == config.SMSProviderBrevo {
		brevoAdapter, err := adapter.NewBrevoAdapter(cfg.Brevo, logger)
		if err != nil {
			return nil, fmt.Errorf("failed to create Brevo SMS adapter: %w", err)
		}
		from = cfg.Brevo.SMSSender
		logger.Info(ctx, "Using Brevo adapter for SMS sending")
		smsAdapter = brevoAdapter
	} else if cfg.Adapter.SMS == config.SMSProviderTwilio {
		twilioAdapter, err := adapter.NewTwilioAdapter(cfg.Twilio, logger)
		if err != nil {
			return nil, fmt.Errorf("failed to create Twilio SMS adapter: %w", err)
		}
		from = cfg.Twilio.FromNumber
		logger.Info(ctx, "Using Twilio adapter for SMS sending")
		smsAdapter = twilioAdapter
	}
	if smsAdapter == nil {
		smsAdapter = adapter.NewMockSMSAdapter(logger)
		logger.Info(ctx, "Using MockSMS adapter for SMS sending")
		from = "MockSender"
	}
	logger.Info(ctx, "SMS service initialized")

	return &smsService{
		adapter: smsAdapter,
		logger: logger.WithFields(map[string]any{
			"service": "sms_service_" + cfg.Adapter.SMS,
		}),
		name: cfg.Adapter.SMS,
		from: from,
	}, nil
}

// Send validates the SMS data and delegates the sending task to the adapter.
func (s *smsService) Send(ctx context.Context, sms dto.SMS) error {
	if sms.To == "" {
		return fmt.Errorf("sms recipient ('To' phone number) cannot be empty")
	}
	if sms.Message == "" {
		return fmt.Errorf("sms message cannot be empty")
	}
	if s.from == "" {
		return fmt.Errorf("sms sender ('From') cannot be empty")
	}

	s.logger.Info(ctx, "Attempting to send SMS via adapter", map[string]any{
		"to":   sms.To,
		"from": s.from,
	})

	// Delegate to the adapter
	err := s.adapter.Send(ctx, sms)
	if err != nil {
		s.logger.Error(ctx, "Failed to send SMS via adapter", map[string]any{"error": err})
		return fmt.Errorf("failed to send SMS via adapter: %w", err)
	}

	s.logger.Info(ctx, "SMS potentially sent successfully via adapter", map[string]any{"to": sms.To})
	return nil
}

// SendCode sends an SMS with a verification code.
func (s *smsService) SendCode(ctx context.Context, to string, code string) error {
	s.logger.Info(ctx, "Sending verification code via SMS", map[string]any{
		"to": to,
	})

	// Create the SMS message
	message := dto.SMS{
		To:      to,
		Message: fmt.Sprintf("Your verification code is: %s. This code will expire in 10 minutes.", code),
	}

	// Send the SMS
	return s.Send(ctx, message)
}

// ServiceName returns the name of the SMS service.
func (s *smsService) ServiceName() string {
	return string(s.name)
}
