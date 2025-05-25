package sen

import (
	"context"

	"github.com/lugondev/send-sen/dto"
)

// Adapter defines the interface for sending emails via different providers.
type EmailAdapter interface {
	SendEmail(ctx context.Context, email dto.Email) error
}

// Service defines the core logic for handling emails.
type EmailService interface {
	SendEmail(ctx context.Context, email dto.Email) error
	SendPasswordReset(ctx context.Context, to string, link string) error
	SendVerificationCode(ctx context.Context, to string, code string) error
	SendWelcome(ctx context.Context, to string, name string) error
	SendWarningLogin(ctx context.Context, to string, location string, time string) error
	ServiceName() string
}
