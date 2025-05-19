package sms_test

import (
	"context"
	"testing"

	"github.com/lugondev/send-sen/config"
	"github.com/lugondev/send-sen/modules/sms/adapter"
	"github.com/lugondev/send-sen/modules/sms/port"
	"github.com/lugondev/send-sen/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestTwilioAdapter_SendSMS(t *testing.T) {
	// Load config and create logger
	cfg, err := config.LoadConfig("../../config")
	assert.NoError(t, err, "Failed to load config")

	log, err := logger.NewZapLogger(cfg)
	assert.NoError(t, err, "Failed to create logger")

	// Create adapter instance
	twilioAdapter, err := adapter.NewTwilioAdapter(cfg.Twilio, log)
	assert.NoError(t, err)

	sms := port.SMS{
		To:      "+18777804236", // Example  number
		Message: "Test SMS from Twilio 123123",
	}

	err = twilioAdapter.SendSMS(context.Background(), sms)
	assert.NoError(t, err)
}
