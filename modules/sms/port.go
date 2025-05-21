package sms

import (
	"context"
	adapter "github.com/lugondev/send-sen/adapters/sms"
)

// Adapter defines the interface for sending SMS messages via different providers.
type Adapter interface {
	Send(ctx context.Context, sms adapter.SMS) error
}

// Service defines the core logic for handling SMS messages.
type Service interface {
	Send(ctx context.Context, sms adapter.SMS) error
	SendCode(ctx context.Context, to string, code string) error
	ServiceName() string
}
