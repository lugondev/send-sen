# Golang Send-Sen

A robust, modular notification service library for Go applications that provides unified interfaces for Email, SMS, and various notification channels. Built with extensibility in mind, it offers a clean architecture with adapters for popular service providers while maintaining flexibility for custom implementations.

## Installation

```bash
go get github.com/lugondev/send-sen
```

## Features

### Email Module ✉️
- Integrated providers:
  - ✅ SendGrid - Enterprise-grade email delivery
  - ✅ Brevo (formerly Sendinblue) - Comprehensive email marketing solution
  - ✅ Mock adapter (for testing) - Simplifies unit testing

### SMS Module 📱
- Integrated providers:
  - ✅ Twilio - Industry-standard SMS service
  - ✅ Brevo SMS - Cost-effective SMS solution
- Extensible adapter interface for custom providers

### Notification Module 🔔
- Integrated providers:
  - ✅ Telegram - Instant messaging platform integration
  - ✅ Mock logging (for testing) - Facilitates testing scenarios
- Flexible port interface for custom providers

## Project Structure

```
.
├── modules/
│   ├── email/         # Email service implementations
│   │   ├── adapter/   # Email provider adapters
│   │   ├── port/      # Email service interfaces
│   │   └── service/   # Email service logic
│   ├── sms/          # SMS service implementations
│   │   ├── adapter/   # SMS provider adapters
│   │   ├── port/      # SMS service interfaces
│   │   └── service/   # SMS service logic
│   └── notify/       # General notification services
│       ├── adapter/   # Notification provider adapters
│       ├── port/      # Notification service interfaces
│       └── service/   # Notification service logic
├── config/          # Configuration management
└── tests/           # Integration tests
```

## Configuration

### YAML Configuration
Create a `config.yaml` file:

```yaml
app:
    name: 'send-sen'
log:
    level: 'debug'    # debug, info, warn, error
    format: 'console' # console, json

sendgrid:
    apiKey: 'your-sendgrid-api-key'
    fromEmail: 'sender@example.com'
    fromName: 'Sender Name'

twilio:
    accountSid: 'your-account-sid'
    messagingSid: 'your-messaging-sid'
    authToken: 'your-auth-token'
    fromNumber: '+1234567890'

brevo:
    apiKey: 'your-brevo-api-key'
    senderEmail: 'sender@example.com'
    senderName: 'Sender Name'
    smsSender: 'SMSSender'

telegram:
    botToken: 'your-bot-token'
    chatId: 'your-chat-id'
    debug: false

adapter:
    notify: 'telegram' # Default notification adapter
    email: 'sendgrid'  # Default email adapter
    sms: 'twilio'      # Default SMS adapter
```

### Code Configuration

```go
import "github.com/lugondev/send-sen/config"

// Load configuration from config.yaml
cfg := config.LoadConfig()

// Or configure programmatically
cfg := &config.Config{
    Email: config.EmailConfig{
        SendGridAPIKey: "your-sendgrid-key",
        BrevoAPIKey:    "your-brevo-key",
    },
    SMS: config.SMSConfig{
        TwilioAccountSID: "your-twilio-sid",
        TwilioAuthToken:  "your-twilio-token",
        BrevoAPIKey:      "your-brevo-key",
    },
    Notify: config.NotifyConfig{
        TelegramBotToken: "your-telegram-token",
    },
}
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

The project includes comprehensive testing support:

### Mock Adapters
- `mock_adapter.go` for email testing
- `mock_adapter.go` for SMS testing
- `mocklog_adapter.go` for notification testing

### Running Tests

```bash
# Run all tests
go test ./tests/...

# Run specific module tests
go test ./tests/email/...
go test ./tests/sms/...
go test ./tests/notify/...

# Run with verbose output
go test -v ./tests/...

# Run with coverage
go test -coverprofile=coverage.out ./tests/...
go tool cover -html=coverage.out
```

### Writing Tests

```go
// Example test using mock adapter
func TestEmailService(t *testing.T) {
    mockAdapter := adapter.NewMockAdapter()
    emailService := service.NewEmailService(mockAdapter)
    
    params := &email.SendParams{
        From:    "test@example.com",
        To:      "recipient@example.com",
        Subject: "Test",
        Body:    "Test message",
    }
    
    err := emailService.Send(context.Background(), params)
    assert.NoError(t, err)
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
- [ ] Microsoft Teams integration

#### Email Providers
- [ ] Mailgun support
- [ ] Mailchimp support
- [ ] Amazon SES integration

#### SMS & Push Notifications
- [ ] Firebase Cloud Messaging (FCM)
- [ ] Vonage (formerly Nexmo)
- [ ] MessageBird

## Contributing

We welcome contributions! Here's how you can help:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Areas for Contribution
- Implement planned integrations
- Improve documentation
- Add more test coverage
- Optimize existing code
- Report bugs
- Suggest new features

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
