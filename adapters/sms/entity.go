package sms

// SMS represents the data structure for an SMS message.
type SMS struct {
	To      string // The recipient's phone number (E.164 format recommended)
	Message string // The text message content
}
