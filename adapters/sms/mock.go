package sms

import (
	"context"
	"fmt"
	"github.com/lugondev/send-sen/modules/sms"

	"github.com/lugondev/send-sen/pkg/logger"
)

// MockSMSAdapter is a mock implementation of the port.EmailAdapter interface.
type MockSMSAdapter struct {
	logger logger.Logger
}

// NewMockSMSAdapter creates a new instance of MockSMSAdapter.
func NewMockSMSAdapter(logger logger.Logger) sms.SMSAdapter {
	namedLogger := logger.WithFields(map[string]any{
		"service": "mock_sms_adapter",
	})
	namedLogger.Info(context.Background(), "Mock SMS adapter initialized")

	return &MockSMSAdapter{
		logger: namedLogger,
	}
}

// SendSMS sends an SMS using the mock implementation.
func (a *MockSMSAdapter) SendSMS(ctx context.Context, sms sms.SMS) error {
	// Placeholder implementation (remove once real implementation is added)
	a.logger.Warn(ctx, "Mock SendSMS function is not fully implemented yet.")
	fmt.Printf("--- MOCK SMS Send ---\nTo: %v\nMessage: %s\n-----------------------\n",
		sms.To, sms.Message)
	// --- End Placeholder ---

	return nil
}
