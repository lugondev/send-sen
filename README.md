# Golang Send-Sen

A comprehensive notification service library for Go applications, providing unified interfaces for Email, SMS, and various notification channels.

## Installation

```bash
go get github.com/lugondev/send-sen
```

## Features

### Email Module ✉️
- Integrated providers:
  - ✅ SendGrid
  - ✅ Brevo (formerly Sendinblue)
  - ✅ Mock adapter (for testing)

### SMS Module 📱
- Integrated providers:
  - ✅ Twilio
  - ✅ Brevo SMS
- Extensible adapter interface for custom providers

### Notification Module 🔔
- Integrated providers:
  - ✅ Telegram
  - ✅ Mock logging (for testing)
- Flexible port interface for custom providers

## Project Structure

```
.
├── modules/
│   ├── email/         # Email service implementations
│   ├── sms/          # SMS service implementations
│   └── notify/       # General notification services
└── tests/            # Integration tests
```

## Usage

### Email Service

```go
import (
    "context"
    "github.com/lugondev/send-sen/modules/email/service"
    "github.com/lugondev/send-sen/modules/email/adapter"
)

// 1. Initialize email adapter (SendGrid example)
emailAdapter := adapter.NewSendGridAdapter(apiKey)

// 2. Create email service
emailService := service.NewEmailService(emailAdapter)

// 3. Prepare email parameters
emailParams := &email.SendParams{
    From:    "sender@example.com",
    To:      "recipient@example.com",
    Subject: "Test Email",
    Body:    "This is a test email",
}

// 4. Send email
err := emailService.Send(context.Background(), emailParams)
if err != nil {
    log.Fatal(err)
}
```

### SMS Service

```go
import (
    "context"
    "github.com/lugondev/send-sen/modules/sms/service"
)

// 1. Configure SMS service
config := &sms.Config{
    APIKey:    "your-api-key",
    APISecret: "your-api-secret",
}

// 2. Initialize SMS service (Brevo example)
smsService := service.NewBrevoService(config)

// 3. Prepare SMS parameters
smsParams := &sms.SendParams{
    From:    "+1234567890",
    To:      "+0987654321",
    Message: "Test SMS message",
}

// 4. Send SMS
err := smsService.Send(context.Background(), smsParams)
if err != nil {
    log.Fatal(err)
}
```

### Notification Service (Telegram)

```go
import (
    "context"
    "github.com/lugondev/send-sen/modules/notify/service"
    "github.com/lugondev/send-sen/modules/notify/adapter"
)

// 1. Initialize Telegram adapter
telegramAdapter := adapter.NewTelegramAdapter(botToken)

// 2. Create notification service
notifyService := service.NewNotifyService(telegramAdapter)

// 3. Prepare notification parameters
notifyParams := &notify.SendParams{
    ChatID:  "your-chat-id",
    Message: "Test notification",
}

// 4. Send notification
err := notifyService.Send(context.Background(), notifyParams)
if err != nil {
    log.Fatal(err)
}
```

## Testing

The project includes mock adapters for testing:
- `mock_adapter.go` for email testing
- `mocklog_adapter.go` for notification testing

Run tests:
```bash
go test ./tests/...
```

## Configuration

Use the `config` package to manage your service configurations:

```go
import "github.com/lugondev/send-sen/config"

// Load configuration
cfg := config.LoadConfig()

// Or configure manually
cfg := &config.Config{
    Email: config.EmailConfig{
        SendGridAPIKey: "your-sendgrid-key",
        BrevoAPIKey:    "your-brevo-key",
    },
    SMS: config.SMSConfig{
        TwilioAccountSID: "your-twilio-sid",
        TwilioAuthToken: "your-twilio-token",
        BrevoAPIKey:     "your-brevo-key",
    },
    Notify: config.NotifyConfig{
        TelegramBotToken: "your-telegram-token",
    },
}
```

## Provider Status

### Integrated Providers ✅
- Email Services:
  - SendGrid - Full support for transactional emails
  - Brevo - Complete email sending capabilities
- SMS Services:
  - Twilio - Full SMS functionality
  - Brevo SMS - Complete SMS support
- Notification Services:
  - Telegram - Complete bot integration

### Planned Integrations 🚀
#### Notification Services
- [ ] Slack integration
- [ ] Discord integration

#### Email Providers
- [ ] Mailgun support
- [ ] Mailchimp support

#### SMS & Push Notifications
- [ ] Firebase Cloud Messaging (FCM) support

## Contributing

Contributions are welcome! Feel free to:
- Implement any of the planned integrations
- Report bugs
- Suggest new features or integrations
- Submit pull requests

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
