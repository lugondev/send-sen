package notify_test

import (
	"context"
	"github.com/lugondev/send-sen/adapters/notify"
	notify2 "github.com/lugondev/send-sen/modules/notify"
	"testing"

	"github.com/lugondev/send-sen/config"
	"github.com/lugondev/send-sen/pkg/logger"
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

	mockLogger, err := logger.NewZapLogger(cfg)
	assert.NoError(t, err)

	telegramCfg := notify.TelegramConfig{
		BotToken: cfg.Telegram.BotToken,
		ChatID:   cfg.Telegram.ChatID,
		Debug:    cfg.Telegram.Debug,
	}

	telegramAdapter, err := notify.NewTelegramAdapter(telegramCfg, mockLogger)
	assert.NoError(t, err)

	// Test with explicit recipient
	err = telegramAdapter.Send(context.Background(), notify2.Content{
		Subject: "Test Subject",
		Message: "Test message with subject from automated test",
	})
	assert.NoError(t, err)
}
