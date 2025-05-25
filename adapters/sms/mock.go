package sms

import (
	"context"
	"fmt"

	logger "github.com/lugondev/go-log"
	"github.com/lugondev/send-sen/dto"
)

// MockSMSAdapter is a mock implementation of the port.EmailAdapter interface.
type MockSMSAdapter struct {
	logger logger.Logger
}

// NewMockSMSAdapter creates a new instance of MockSMSAdapter.
func NewMockSMSAdapter(logger logger.Logger) *MockSMSAdapter {
	namedLogger := logger.WithFields(map[string]any{
		"service": "mock_sms_adapter",
	})
	namedLogger.Info(context.Background(), "Mock SMS adapter initialized")

	return &MockSMSAdapter{
		logger: namedLogger,
	}
}

// Send sends an SMS using the mock implementation.
func (a *MockSMSAdapter) Send(ctx context.Context, sms dto.SMS) error {
	// Placeholder implementation (remove once real implementation is added)
	a.logger.Warn(ctx, "Mock SendSMS function is not fully implemented yet.")
	fmt.Printf("--- MOCK SMS Send ---\nTo: %v\nMessage: %s\n-----------------------\n",
		sms.To, sms.Message)
	// --- End Placeholder ---

	return nil
}
