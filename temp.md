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
