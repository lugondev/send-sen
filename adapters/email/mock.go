package email

import (
	"context"
	"fmt"

	"github.com/lugondev/go-log"
)

// MockEmailAdapter is a mock implementation of the port.EmailAdapter interface.
type MockEmailAdapter struct {
	logger logger.Logger
}

// NewMockEmailAdapter creates a new instance of MockEmailAdapter.
func NewMockEmailAdapter(logger logger.Logger) *MockEmailAdapter {
	namedLogger := logger.WithFields(map[string]any{
		"service": "mock_email",
	})
	adapter := &MockEmailAdapter{
		logger: namedLogger,
	}
	namedLogger.Info(context.Background(), "Mock email adapter initialized")

	return adapter
}

// SendEmail sends an email using the mock implementation.
func (a *MockEmailAdapter) SendEmail(ctx context.Context, email Email) error {
	// Placeholder implementation (remove once real implementation is added)
	a.logger.Warn(ctx, "Mock SendEmail function is not fully implemented yet.")
	fmt.Printf("--- MOCK Brevo Send ---\nTo: %v\nCc: %v\nBcc: %v\nSubject: %s\nBody: %s\n-----------------------\n",
		email.To, email.Cc, email.Bcc, email.Subject, email.Body)
	// --- End Placeholder ---

	return nil // Return nil for now
}
