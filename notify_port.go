package main

import (
	"context"

	"github.com/lugondev/send-sen/dto"
)

// NotifyAdapter defines the interface for sending notifications via different providers/channels.
type NotifyAdapter interface {
	Send(ctx context.Context, content dto.Content) error
}

// NotifyService defines the core logic for handling notifications.
type NotifyService interface {
	Send(ctx context.Context, content dto.Content) error
	Alert(ctx context.Context, subject, message string) error
	Info(ctx context.Context, subject, message string) error
	Notify(ctx context.Context, subject, message string, level dto.Level) error
	ServiceName() string
}
