package port

import "context"

// Email represents the email data structure.
type Email struct {
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Html    string
	Body    string
}

// EmailAdapter defines the interface for sending emails via different providers.
type EmailAdapter interface {
	SendEmail(ctx context.Context, email Email) error
}

// EmailService defines the core logic for handling emails.
type EmailService interface {
	SendEmail(ctx context.Context, email Email) error
	ServiceName() string
}
