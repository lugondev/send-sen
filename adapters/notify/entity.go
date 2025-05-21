package notify

type Level string

const (
	Debug   Level = "debug"
	Info    Level = "info"
	Warning Level = "warning"
	Error   Level = "error"
)

// Content represents the data structure for a generic notification.
// We might need more specific structures or fields depending on the channel.
type Content struct {
	Subject   string
	Message   string
	Level     Level
	ParseMode string
}
