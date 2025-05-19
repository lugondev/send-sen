package port

import "context"

// SMS represents the data structure for an SMS message.
type SMS struct {
	To      string // The recipient's phone number (E.164 format recommended)
	Message string // The text message content
}

// SMSAdapter defines the interface for sending SMS messages via different providers.
type SMSAdapter interface {
	SendSMS(ctx context.Context, sms SMS) error
}

// SMSService defines the core logic for handling SMS messages.
type SMSService interface {
	SendSMS(ctx context.Context, sms SMS) error
	ServiceName() string
}
