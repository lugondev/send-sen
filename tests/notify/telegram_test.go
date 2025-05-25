package notify_test

import (
	"context"
	"testing"

	logger "github.com/lugondev/go-log"
	"github.com/lugondev/send-sen/adapters/notify"
	"github.com/lugondev/send-sen/config"
	"github.com/lugondev/send-sen/domain/dto"
	"github.com/stretchr/testify/assert"
)

func TestTelegramSendNotification_RealConfig(t *testing.T) {
	t.Log("Testing Telegram notification with real config...")

	// Skip if running in CI/CD
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	cfg, err := config.LoadConfig("../../config")
	assert.NoError(t, err)

	mockLogger, err := logger.NewLogger(&logger.Option{
		Format:       cfg.Log.Format,
		ScopeName:    "send-sen",
		ScopeVersion: "v0.1.1",
	})
	assert.NoError(t, err)

	telegramAdapter, err := notify.NewTelegramAdapter(cfg.Telegram, mockLogger)
	assert.NoError(t, err)

	// Test with explicit recipient
	err = telegramAdapter.Send(context.Background(), dto.Content{
		Subject: "Test Subject",
		Message: "Test message with subject from automated test",
	})
	assert.NoError(t, err)

	// Test with level and parse mode
	err = telegramAdapter.Send(context.Background(), dto.Content{
		Subject:   "Error Alert",
		Message:   "This is a test error message with formatting",
		Level:     dto.Error,
		ParseMode: "HTML",
	})
	assert.NoError(t, err)

	// Test with warning level
	err = telegramAdapter.Send(context.Background(), dto.Content{
		Subject: "Warning Notice",
		Message: "This is a test warning message",
		Level:   dto.Warning,
	})
	assert.NoError(t, err)
}
