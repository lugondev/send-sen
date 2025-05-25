package notify

import (
	"context"
	"fmt"

	logger "github.com/lugondev/go-log"
	"github.com/lugondev/send-sen/domain/dto"
)

// MockLogAdapter implements the port.NotifyAdapter interface by logging notifications.
type MockLogAdapter struct {
	logger logger.Logger
}

// NewMockLogAdapter creates a new instance of MockLogAdapter.
func NewMockLogAdapter(logger logger.Logger) *MockLogAdapter {
	namedLogger := logger.WithFields(map[string]any{
		"service": "mocklog_notify_adapter",
	})
	ctx := context.Background()
	namedLogger.Info(ctx, "MockLog notify adapter initialized")

	return &MockLogAdapter{logger: namedLogger}
}

// Send logs the notification details.
func (a *MockLogAdapter) Send(ctx context.Context, msg dto.Content) error {
	// Get level-specific icon
	var levelIcon string
	switch msg.Level {
	case dto.Debug:
		levelIcon = "[DEBUG]"
	case dto.Info:
		levelIcon = "[INFO]"
	case dto.Warning:
		levelIcon = "[WARNING]"
	case dto.Error:
		levelIcon = "[ERROR]"
	default:
		levelIcon = "[NOTIFICATION]"
	}

	// Format the message for logging
	var formattedMessage string
	if msg.Subject != "" {
		formattedMessage = fmt.Sprintf("%s %s\n%s", levelIcon, msg.Subject, msg.Message)
	} else {
		formattedMessage = fmt.Sprintf("%s %s", levelIcon, msg.Message)
	}

	// Log with appropriate level
	logFields := map[string]any{
		"level": msg.Level,
	}

	if msg.Subject != "" {
		logFields["subject"] = msg.Subject
	}

	a.logger.Info(ctx, "--- MOCK Notification Sent (via Log) ---", map[string]any{
		"formatted_message": formattedMessage,
		"level":             string(msg.Level),
	})
	return nil
}
