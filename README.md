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

## Architecture

The project follows a clean architecture pattern:
1. **Modules**: Contain the core business logic and define interfaces (ports)
2. **Adapters**: Implement the interfaces defined by modules to connect with external services
3. **Configuration**: Centralized configuration management using Viper
4. **Testing**: Comprehensive test suite with mock adapters for testing

## Project Structure

```
.
├── adapters/               # Implementation of service adapters
│   ├── email/              # Email provider adapters (SendGrid, Brevo, Mock)
│   ├── notify/             # Notification provider adapters (Telegram, Mock)
│   └── sms/                # SMS provider adapters (Twilio, Brevo, Mock)
├── config/                 # Configuration management
│   ├── config.go           # Configuration structures and loading logic
│   ├── config.example.yaml # Example configuration file
│   └── config.yaml         # Actual configuration file (gitignored)
├── modules/                # Core business logic modules
│   ├── email/              # Email service module
│   ├── notify/             # Notification service module
│   └── sms/                # SMS service module
├── pkg/                    # Shared packages
│   └── logger/             # Logging functionality
└── tests/                  # Integration tests
    ├── email/              # Email service tests
    ├── notify/             # Notification service tests
    └── sms/                # SMS service tests
```

## Configuration

### YAML Configuration
Create a `config.yaml` file:

## Testing
The project includes comprehensive testing support:

### Mock Adapters
- Mock adapters in `adapters/email/mock.go` for email testing
- Mock adapters in `adapters/sms/mock.go` for SMS testing
- Mock adapters in `adapters/notify/mock.go` for notification testing

### Running Tests

```bash
# Run all tests
go test ./tests/...

# Run specific module tests
go test ./tests/email/...
go test ./tests/sms/...
go test ./tests/notify/...

# Check test coverage for critical changes
go test -coverprofile=coverage.out ./tests/...
go tool cover -html=coverage.out
```

## Code Style Guidelines

1. Follow Go's standard code style and conventions
2. Use meaningful variable and function names
3. Add appropriate comments for public functions and complex logic
4. Implement proper error handling and logging
5. Keep functions small and focused on a single responsibility
6. Write unit tests for new functionality

## Implementation Guidelines

### For new adapters
- Implement the appropriate interface from the modules package
- Add comprehensive logging
- Include proper error handling
- Add tests in the tests directory

### For bug fixes
- Identify the root cause
- Add tests that reproduce the issue
- Fix the issue
- Verify that the tests pass

### For refactoring
- Ensure all tests pass before and after changes
- Maintain backward compatibility when possible
- Update documentation if interfaces change

## Building and Running

The project is a library, so there's no need to build it separately. However, when testing changes:

```bash
# Verify that the code compiles
go build ./...

# Run tests
go test ./tests/...
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

## Examples & Tutorials

### Email Module ✉️

The Email module provides a simple way to send emails using different providers.

#### Usage

First, you need to initialize the Email service with a specific provider. Here's an example using Brevo:

```go
package main

import (
	"context"
	"log"

	"github.com/lugondev/send-sen/adapters/email"
	"github.com/lugondev/send-sen/config"
	emailModule "github.com/lugondev/send-sen/modules/email"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize Brevo adapter
	brevoAdapter := email.NewBrevoAdapter(cfg.Brevo.APIKey)

	// Initialize Email service
	emailService := emailModule.NewEmailService(brevoAdapter)

	// Create email data
	emailData := emailModule.EmailData{
		To:      []string{"recipient@example.com"},
		Subject: "Hello from Send-Sen",
		Body:    "This is a test email sent using Send-Sen.",
	}

	// Send email
	err = emailService.Send(context.Background(), emailData)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	log.Println("Email sent successfully!")
}
```

#### Configuration

You need to configure the Brevo API key in your `config.yaml` file:

```yaml
brevo:
  api_key: YOUR_BREVO_API_KEY
```

### SMS Module 📱

The SMS module allows you to send SMS messages using various providers.

#### Usage

Here's an example using Twilio:

```go
package main

import (
	"context"
	"log"

	"github.com/lugondev/send-sen/adapters/sms"
	"github.com/lugondev/send-sen/config"
	smsModule "github.com/lugondev/send-sen/modules/sms"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize Twilio adapter
	twilioAdapter := sms.NewTwilioAdapter(cfg.Twilio.AccountSID, cfg.Twilio.AuthToken, cfg.Twilio.FromNumber)

	// Initialize SMS service
	smsService := smsModule.NewSMSService(twilioAdapter)

	// Create SMS data
	smsData := smsModule.SMSData{
		To:   "+1234567890",
		Body: "Hello from Send-Sen!",
	}

	// Send SMS
	err = smsService.Send(context.Background(), smsData)
	if err != nil {
		log.Fatalf("Failed to send SMS: %v", err)
	}

	log.Println("SMS sent successfully!")
}
```

#### Configuration

Configure the Twilio credentials in your `config.yaml` file:

```yaml
twilio:
  account_sid: YOUR_TWILIO_ACCOUNT_SID
  auth_token: YOUR_TWILIO_AUTH_TOKEN
  from_number: YOUR_TWILIO_FROM_NUMBER
```

### Notification Module 🔔

The Notification module provides a way to send notifications via different channels, such as Telegram.

#### Usage

Here's an example using Telegram:

```go
package main

import (
	"context"
	"log"

	"github.com/lugondev/send-sen/adapters/notify"
	"github.com/lugondev/send-sen/config"
	notifyModule "github.com/lugondev/send-sen/modules/notify"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize Telegram adapter
	telegramAdapter := notify.NewTelegramAdapter(cfg.Telegram.BotToken, cfg.Telegram.ChatID)

	// Initialize Notify service
	notifyService := notifyModule.NewNotifyService(telegramAdapter)

	// Create notification data
	notificationData := notifyModule.NotificationData{
		Message: "Hello from Send-Sen!",
	}

	// Send notification
	err = notifyService.Send(context.Background(), notificationData)
	if err != nil {
		log.Fatalf("Failed to send notification: %v", err)
	}

	log.Println("Notification sent successfully!")
}
```

#### Configuration

Configure the Telegram bot token and chat ID in your `config.yaml` file:

```yaml
telegram:
  bot_token: YOUR_TELEGRAM_BOT_TOKEN
  chat_id: YOUR_TELEGRAM_CHAT_ID
```
