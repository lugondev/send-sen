package service

import (
	"context"
	"fmt"

	"github.com/lugondev/send-sen/config"
	"github.com/lugondev/send-sen/modules/sms/adapter"
	"github.com/lugondev/send-sen/modules/sms/port"
	"github.com/lugondev/send-sen/pkg/logger"
)

// smsService implements the SMSService interface.
type smsService struct {
	adapter port.SMSAdapter
	logger  logger.Logger
	name    config.SMSProvider
	from    string
}

// NewSMSService creates a new instance of SMSService.
// It requires an SMSAdapter to be provided.
func NewSMSService(cfg config.Config, logger logger.Logger) (port.SMSService, error) {
	namedLogger := logger.WithFields(map[string]any{
		"service": "sms_service_" + cfg.Adapter.SMS,
	})
	ctx := context.Background()
	var smsAdapter port.SMSAdapter
	var from string
	if cfg.Adapter.SMS == config.SMSProviderBrevo {
		brevoAdapter, err := adapter.NewBrevoAdapter(cfg.Brevo, namedLogger)
		if err != nil {
			return nil, fmt.Errorf("failed to create Brevo SMS adapter: %w", err)
		}
		from = cfg.Brevo.SMSSender
		namedLogger.Info(ctx, "Using Brevo adapter for SMS sending")
		smsAdapter = brevoAdapter
	} else if cfg.Adapter.SMS == config.SMSProviderTwilio {
		twilioAdapter, err := adapter.NewTwilioAdapter(cfg.Twilio, namedLogger)
		if err != nil {
			return nil, fmt.Errorf("failed to create Twilio SMS adapter: %w", err)
		}
		from = cfg.Twilio.FromNumber
		namedLogger.Info(ctx, "Using Twilio adapter for SMS sending")
		smsAdapter = twilioAdapter
	}
	if smsAdapter == nil {
		smsAdapter = adapter.NewMockSMSAdapter(namedLogger)
		namedLogger.Info(ctx, "Using MockSMS adapter for SMS sending")
		from = "MockSender"
	}
	namedLogger.Info(ctx, "SMS service initialized")

	return &smsService{
		adapter: smsAdapter,
		logger:  namedLogger,
		name:    cfg.Adapter.SMS,
		from:    from,
	}, nil
}

// SendSMS validates the SMS data and delegates the sending task to the adapter.
func (s *smsService) SendSMS(ctx context.Context, sms port.SMS) error {
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
	err := s.adapter.SendSMS(ctx, sms)
	if err != nil {
		s.logger.Error(ctx, "Failed to send SMS via adapter", map[string]any{"error": err})
		return fmt.Errorf("failed to send SMS via adapter: %w", err)
	}

	s.logger.Info(ctx, "SMS potentially sent successfully via adapter", map[string]any{"to": sms.To})
	return nil
}

// ServiceName returns the name of the SMS service.
func (s *smsService) ServiceName() string {
	return string(s.name)
}
