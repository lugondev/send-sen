package email

import (
	"context"
	adapter "github.com/lugondev/send-sen/adapters/email"
)

// Adapter defines the interface for sending emails via different providers.
type Adapter interface {
	SendEmail(ctx context.Context, email adapter.Email) error
}

// Service defines the core logic for handling emails.
type Service interface {
	SendEmail(ctx context.Context, email adapter.Email) error
	SendPasswordReset(ctx context.Context, to string, link string) error
	SendVerificationCode(ctx context.Context, to string, code string) error
	SendWelcome(ctx context.Context, to string, name string) error
	SendWarningLogin(ctx context.Context, to string, location string, time string) error
	ServiceName() string
}
