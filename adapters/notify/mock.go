package notify

import (
	"context"
	"github.com/lugondev/send-sen/modules/notify"

	"github.com/lugondev/send-sen/pkg/logger"
)

// MockLogAdapter implements the port.NotifyAdapter interface by logging notifications.
type MockLogAdapter struct {
	logger logger.Logger
}

// NewMockLogAdapter creates a new instance of MockLogAdapter.
func NewMockLogAdapter(logger logger.Logger) notify.NotifyAdapter {
	namedLogger := logger.WithFields(map[string]any{
		"service": "mocklog_notify_adapter",
	})
	ctx := context.Background()
	namedLogger.Info(ctx, "MockLog notify adapter initialized")

	return &MockLogAdapter{logger: namedLogger}
}

// Send logs the notification details.
func (a *MockLogAdapter) Send(ctx context.Context, notification notify.Content) error {
	a.logger.Info(ctx, "--- MOCK Notification Sent (via Log) ---", map[string]any{
		"subject": notification.Subject,
		"message": notification.Message,
	})
	return nil
}
