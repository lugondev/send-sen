package email

import (
	"context"
	"testing"

	"github.com/lugondev/send-sen/config"
	"github.com/lugondev/send-sen/modules/email/adapter"
	"github.com/lugondev/send-sen/modules/email/port"
	"github.com/lugondev/send-sen/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestNewBrevoAdapter(t *testing.T) {
	// Load real config
	cfg, err := config.LoadConfig("../../config")
	assert.NoError(t, err, "Failed to load config")

	// Initialize logger
	log, err := logger.NewZapLogger(cfg)
	assert.NoError(t, err, "Failed to create logger")

	brevoAdapter, err := adapter.NewBrevoAdapter(cfg.Brevo, log)
	assert.NoError(t, err, "Failed to create Brevo adapter")
	assert.NotNil(t, brevoAdapter, "Brevo adapter should not be nil")

	err = brevoAdapter.SendEmail(context.Background(), port.Email{
		To:      []string{"lugondev@gmail.com"},
		Subject: "Test Email",
		Body:    "This is a test email",
		Html:    "<p>This is a test email</p>",
	})
	assert.NoError(t, err, "Failed to send email")
}
