package ports

import (
	"context"

	"github.com/lugondev/send-sen/domain/dto"
)

// SMSAdapter defines the interface for sending SMS messages via different providers.
type SMSAdapter interface {
	Send(ctx context.Context, sms dto.SMS) error
}

// SMSService defines the core logic for handling SMS messages.
type SMSService interface {
	Send(ctx context.Context, sms dto.SMS) error
	SendCode(ctx context.Context, to string, code string) error
	ServiceName() string
}
