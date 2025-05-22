package email

import (
	"context"
	"github.com/lugondev/send-sen/adapters/email"
	"testing"

	"github.com/lugondev/go-log"
	"github.com/lugondev/send-sen/config"
	"github.com/stretchr/testify/assert"
)

func TestNewBrevoAdapter(t *testing.T) {
	// Load real config
	cfg, err := config.LoadConfig("../../config")
	assert.NoError(t, err, "Failed to load config")

	// Initialize logger
	log, err := logger.NewLogger(&logger.Option{
		Format:       cfg.Log.Format,
		ScopeName:    "send-sen",
		ScopeVersion: "v0.1.1",
	})
	assert.NoError(t, err, "Failed to create logger")

	brevoAdapter, err := email.NewBrevoAdapter(cfg.Brevo, log)
	assert.NoError(t, err, "Failed to create Brevo adapter")
	assert.NotNil(t, brevoAdapter, "Brevo adapter should not be nil")

	// Create a message using the adapter's Email type
	err = brevoAdapter.SendEmail(context.Background(), email.Email{
		To:      []string{"lugondev@gmail.com"},
		Subject: "Test Message",
		Body:    "This is a test email",
		Html:    "<p>This is a test email</p>",
	})
	assert.NoError(t, err, "Failed to send email")
}
