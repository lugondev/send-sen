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

func TestBrevoAdapter_SendSMS(t *testing.T) {
	// Load config and create logger
	cfg, err := config.LoadConfig("../../config")
	assert.NoError(t, err, "Failed to load config")

	log, err := logger.NewZapLogger(cfg)
	assert.NoError(t, err, "Failed to create logger")

	// Create adapter instance
	brevoAdapter, err := adapter.NewBrevoAdapter(cfg.Brevo, log)
	assert.NoError(t, err)

	sms := port.SMS{
		To:      "+84909123456", // Example Vietnamese number
		Message: "Test SMS from Brevo 123123",
	}

	err = brevoAdapter.SendSMS(context.Background(), sms)
	assert.NoError(t, err)
}
