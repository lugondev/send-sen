package email

type Email struct {
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Html    string
	Body    string
}
