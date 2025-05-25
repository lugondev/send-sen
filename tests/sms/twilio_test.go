package sms_test

import (
	"context"
	"testing"

	logger "github.com/lugondev/go-log"
	adapter "github.com/lugondev/send-sen/adapters/sms"
	"github.com/lugondev/send-sen/domain/dto"

	"github.com/lugondev/send-sen/config"
	"github.com/stretchr/testify/assert"
)

func TestTwilioAdapter_SendSMS(t *testing.T) {
	// Load config and create logger
	cfg, err := config.LoadConfig("../../config")
	assert.NoError(t, err, "Failed to load config")

	log, err := logger.NewLogger(&logger.Option{
		Format:       cfg.Log.Format,
		ScopeName:    "send-sen",
		ScopeVersion: "v0.1.1",
	})
	assert.NoError(t, err, "Failed to create logger")

	// Create adapter instance
	twilioAdapter, err := adapter.NewTwilioAdapter(cfg.Twilio, log)
	assert.NoError(t, err)

	sms := dto.SMS{
		To:      "+18777804236", // Example number
		Message: "Test SMS from Twilio 123123",
	}

	err = twilioAdapter.Send(context.Background(), sms)
	assert.NoError(t, err)
}
