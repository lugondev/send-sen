package sms_test

import (
	"context"
	"github.com/lugondev/send-sen/adapters/sms"
	sms2 "github.com/lugondev/send-sen/modules/sms"
	"testing"

	"github.com/lugondev/send-sen/config"
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
	twilioAdapter, err := sms.NewTwilioAdapter(cfg.Twilio, log)
	assert.NoError(t, err)

	sms := sms2.SMS{
		To:      "+18777804236", // Example  number
		Message: "Test SMS from Twilio 123123",
	}

	err = twilioAdapter.SendSMS(context.Background(), sms)
	assert.NoError(t, err)
}
