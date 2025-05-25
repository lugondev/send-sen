package sms_test

import (
	"context"
	"testing"

	adapter "github.com/lugondev/send-sen/adapters/sms"
	"github.com/lugondev/send-sen/dto"

	logger "github.com/lugondev/go-log"
	"github.com/lugondev/send-sen/config"
	"github.com/stretchr/testify/assert"
)

func TestBrevoAdapter_SendSMS(t *testing.T) {
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
	brevoAdapter, err := adapter.NewBrevoAdapter(cfg.Brevo, log)
	assert.NoError(t, err)

	sms := dto.SMS{
		To:      "+84909123456", // Example Vietnamese number
		Message: "Test SMS from Brevo 123123",
	}

	err = brevoAdapter.Send(context.Background(), sms)
	assert.NoError(t, err)
}
