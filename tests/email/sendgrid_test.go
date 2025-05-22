package email

import (
	"context"
	"github.com/lugondev/send-sen/adapters/email"
	"testing"

	"github.com/lugondev/go-log"
	"github.com/lugondev/send-sen/config"
	"github.com/stretchr/testify/assert"
)

func TestNewSendgridAdapter(t *testing.T) {
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

	sendgridAdapter, err := email.NewSendGridAdapter(cfg.SendGrid, log)
	assert.NoError(t, err, "Failed to create SendGrid adapter")
	assert.NotNil(t, sendgridAdapter, "SendGrid adapter should not be nil")

	err = sendgridAdapter.SendEmail(context.Background(), email.Email{
		To:      []string{"lugondev@gmail.com"},
		Subject: "Test Message",
		Body:    "This is a test email",
		Html:    "<p>This is a test email</p>",
	})
	assert.NoError(t, err, "Failed to send email")
}
