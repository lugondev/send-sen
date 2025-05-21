package notify

import (
	"context"
	adapter "github.com/lugondev/send-sen/adapters/notify"
)

// Adapter defines the interface for sending notifications via different providers/channels.
type Adapter interface {
	Send(ctx context.Context, content adapter.Content) error
}

// Service defines the core logic for handling notifications.
type Service interface {
	Send(ctx context.Context, content adapter.Content) error
	Alert(ctx context.Context, subject, message string) error
	Info(ctx context.Context, subject, message string) error
	Notify(ctx context.Context, subject, message string, level adapter.Level) error
	ServiceName() string
}
