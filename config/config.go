package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// AppConfig stores application-specific configuration.
type AppConfig struct {
	Name string `mapstructure:"name"`
}

type NotifyChannel string

const (
	NotifyMock     NotifyChannel = "mock"
	NotifyTelegram NotifyChannel = "telegram"
	NotifySlack    NotifyChannel = "slack"
)

type EmailProvider string

const (
	EmailSendGrid EmailProvider = "sendgrid"
	EmailBrevo    EmailProvider = "brevo"
	EmailMock     EmailProvider = "mock"
)

type SMSProvider string

const (
	SMSProviderTwilio SMSProvider = "twilio"
	SMSProviderBrevo  SMSProvider = "brevo"
	SMSProviderMock   SMSProvider = "mock"
)

// AdapterConfig holds configuration for different notification adapters.
type AdapterConfig struct {
	Notify NotifyChannel `mapstructure:"notify"`
	Email  EmailProvider `mapstructure:"email"`
	SMS    SMSProvider   `mapstructure:"sms"`
}

// LogConfig stores logging-specific configuration.
type LogConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

// Config stores all configuration of the application.
type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Log      LogConfig      `mapstructure:"log"`
	Adapter  AdapterConfig  `mapstructure:"adapter"`
	SendGrid SendGridConfig `mapstructure:"sendgrid"`
	Twilio   TwilioConfig   `mapstructure:"twilio"`
	Telegram TelegramConfig `mapstructure:"telegram"`
	Brevo    BrevoConfig    `mapstructure:"brevo"`
}

// SendGridConfig holds SendGrid specific configuration.
type SendGridConfig struct {
	APIKey    string `mapstructure:"apiKey"`
	FromEmail string `mapstructure:"fromEmail"`
	FromName  string `mapstructure:"fromName"`
}

// TwilioConfig holds Twilio specific configuration.
type TwilioConfig struct {
	AccountSid   string `mapstructure:"accountSid"`
	AuthToken    string `mapstructure:"authToken"`
	FromNumber   string `mapstructure:"fromNumber"`
	MessagingSid string `mapstructure:"messagingSid"`
}

// TelegramConfig holds Telegram specific configuration.
type TelegramConfig struct {
	BotToken string `mapstructure:"botToken"`
	ChatID   string `mapstructure:"chatId"`
	Debug    bool   `mapstructure:"debug"`
}

// BrevoConfig holds Brevo (formerly Sendinblue) specific configuration.
type BrevoConfig struct {
	APIKey      string `mapstructure:"apiKey"`
	SenderEmail string `mapstructure:"senderEmail"`
	SenderName  string `mapstructure:"senderName"`
	SMSSender   string `mapstructure:"smsSender"`
}

// LoadConfig reads configuration from a YAML file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config") // Look for config.yaml
	viper.SetConfigType("yaml")   // Specify YAML format

	// Attempt to read the config file first
	if readErr := viper.ReadInConfig(); readErr != nil {
		log.Fatal("Error reading config file:", readErr)
	} else {
		log.Println("Using configuration file:", viper.ConfigFileUsed())
	}

	// Allow overriding config values with environment variables
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // Replace dots with underscores for nested keys in env vars (e.g., FIREBASE.PROJECT_ID -> FIREBASE_PROJECT_ID)

	// Unmarshal the config into the struct using the mapstructure tags
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("failed to unmarshal config: %x", err)
	}

	return config, nil
}
