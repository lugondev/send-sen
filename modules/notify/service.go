package notify

import (
	"context"
	"fmt"
	adapter "github.com/lugondev/send-sen/adapters/notify"

	"github.com/lugondev/send-sen/config"
	"github.com/lugondev/send-sen/pkg/logger"
)

// notifyService implements the Service interface.
// It holds a map of registered adapters keyed by channel name.
type notifyService struct {
	adapter Adapter
	logger  logger.Logger
	name    config.NotifyChannel
}

// NewNotifyService creates a new instance of Service.
func NewNotifyService(cfg config.Config, logger logger.Logger) (Service, error) {
	namedLogger := logger.WithFields(map[string]any{
		"service": "notify_service_" + cfg.Adapter.Notify,
	})

	ctx := context.Background()
	logger.Debug(ctx, "Registered notify adapter", map[string]any{
		"channel": cfg.Adapter.Notify,
	})

	var notifyAdapter Adapter
	if cfg.Adapter.Notify == config.NotifyTelegram {
		telegramAdapter, err := adapter.NewTelegramAdapter(cfg.Telegram, namedLogger)
		if err != nil {
			namedLogger.Error(ctx, "Failed to create Telegram adapter", map[string]any{
				"error": err,
			})
			return nil, fmt.Errorf("failed to create Telegram adapter: %w", err)
		} else {
			notifyAdapter = telegramAdapter
			namedLogger.Info(ctx, "Using Telegram adapter for notifications", map[string]any{
				"chat_id": cfg.Telegram.ChatID,
			})
		}
	}
	if notifyAdapter == nil {
		notifyAdapter = adapter.NewMockLogAdapter(namedLogger)
		namedLogger.Info(ctx, "Using MockLog adapter for notifications", map[string]any{
			"channel": cfg.Adapter.Notify,
		})
	}

	return &notifyService{
		adapter: notifyAdapter,
		logger:  namedLogger,
		name:    cfg.Adapter.Notify,
	}, nil
}

// Send finds the appropriate adapter based on the notification's channel
func (s *notifyService) Send(ctx context.Context, content adapter.Content) error {
	if content.Message == "" {
		return fmt.Errorf("notification message cannot be empty")
	}

	s.logger.Info(ctx, "Sending notification via adapter", map[string]any{
		"sub": content.Subject, // Subject might be empty
		"msg": content.Message,
	})

	err := s.adapter.Send(ctx, content)
	if err != nil {
		s.logger.Error(ctx, "Failed to send notification", map[string]any{
			"error": err,
		})
		return fmt.Errorf("failed to send notification: %w", err)
	}

	s.logger.Info(ctx, "Notification sent successfully")
	return nil
}

// Alert sends a notification with Error level
func (s *notifyService) Alert(ctx context.Context, subject, message string) error {
	content := adapter.Content{
		Subject: subject,
		Message: message,
		Level:   adapter.Error,
	}
	return s.Send(ctx, content)
}

// Info sends a notification with Info level
func (s *notifyService) Info(ctx context.Context, subject, message string) error {
	content := adapter.Content{
		Subject: subject,
		Message: message,
		Level:   adapter.Info,
	}
	return s.Send(ctx, content)
}

// Notify sends a notification with the specified level
func (s *notifyService) Notify(ctx context.Context, subject, message string, level adapter.Level) error {
	content := adapter.Content{
		Subject: subject,
		Message: message,
		Level:   level,
	}
	return s.Send(ctx, content)
}

// ServiceName returns the name of the notification service.
func (s *notifyService) ServiceName() string {
	return string(s.name)
}
