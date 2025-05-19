package service

import (
	"context"
	"fmt"

	"github.com/lugondev/send-sen/modules/sms/port"
	"github.com/lugondev/send-sen/pkg/logger"
)

// smsService implements the SMSService interface.
type smsService struct {
	adapter port.SMSAdapter
	logger  logger.Logger
}

// NewSMSService creates a new instance of SMSService.
// It requires an SMSAdapter to be provided.
func NewSMSService(adapter port.SMSAdapter, logger logger.Logger) (port.SMSService, error) {
	if adapter == nil {
		return nil, fmt.Errorf("SMSAdapter cannot be nil")
	}
	namedLogger := logger.WithFields(map[string]any{
		"service": "sms_service",
	})
	ctx := context.Background()
	namedLogger.Info(ctx, "SMS service initialized")
	return &smsService{
		adapter: adapter,
		logger:  namedLogger,
	}, nil
}

// SendSMS validates the SMS data and delegates the sending task to the adapter.
func (s *smsService) SendSMS(sms port.SMS) error {
	if sms.To == "" {
		return fmt.Errorf("sms recipient ('To' phone number) cannot be empty")
	}
	if sms.Message == "" {
		return fmt.Errorf("sms message cannot be empty")
	}
	// 'From' might be optional depending on the provider/adapter configuration
	// if sms.From == "" {
	// 	 return fmt.Errorf("sms sender ('From') cannot be empty")
	// }

	s.logger.Info(context.Background(), "Attempting to send SMS via adapter", map[string]any{
		"to":   sms.To,
		"from": sms.From,
	})

	// Delegate to the adapter
	err := s.adapter.SendSMS(sms)
	if err != nil {
		s.logger.Error(context.Background(), "Failed to send SMS via adapter", map[string]any{"error": err})
		return fmt.Errorf("failed to send SMS via adapter: %w", err)
	}

	s.logger.Info(context.Background(), "SMS potentially sent successfully via adapter", map[string]any{"to": sms.To})
	return nil
}

// TODO: Implement SendBatchSMS if needed.
// func (s *smsService) SendBatchSMS(smsList []port.SMS) []error {
// 	 // Implementation details...
// 	 return nil
// }
