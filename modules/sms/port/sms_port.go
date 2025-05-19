package port

// SMS represents the data structure for an SMS message.
type SMS struct {
	To      string // The recipient's phone number (E.164 format recommended)
	From    string // The sender ID or phone number (provider-specific constraints may apply)
	Message string // The text message content
	// Add other provider-specific fields if necessary (e.g., Type, StatusCallback)
	// Type string // e.g., "Transactional", "Marketing"
}

// SMSAdapter defines the interface for sending SMS messages via different providers.
type SMSAdapter interface {
	SendSMS(sms SMS) error
	// HealthCheck() error // Optional
}

// SMSService defines the core logic for handling SMS messages.
type SMSService interface {
	SendSMS(sms SMS) error
	// SendBatchSMS(smsList []SMS) []error // Optional
}
