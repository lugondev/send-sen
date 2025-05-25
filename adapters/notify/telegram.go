package notify

import (
	"context"
	"fmt"
	"strconv"

	"github.com/lugondev/send-sen/config"
	"github.com/lugondev/send-sen/dto"
	"golang.org/x/net/html"

	logger "github.com/lugondev/go-log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// TelegramAdapter implements the port.NotifyAdapter
type TelegramAdapter struct {
	bot         *tgbotapi.BotAPI
	chatID      int64
	logger      logger.Logger
	serviceName string
}

// NewTelegramAdapter creates a new instance of TelegramAdapter.
func NewTelegramAdapter(config config.TelegramConfig, logger logger.Logger) (*TelegramAdapter, error) {
	if config.BotToken == "" {
		return nil, fmt.Errorf("telegram bot token is required")
	}
	namedLogger := logger.WithFields(map[string]any{
		"service": "telegram_notify_adapter",
	})
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		namedLogger.Error(context.Background(), "Error creating Telegram bot API instance", map[string]any{"error": err})
		return nil, fmt.Errorf("failed to initialize Telegram bot: %w", err)
	}
	bot.Debug = config.Debug
	adapter := &TelegramAdapter{
		bot:         bot,
		logger:      namedLogger,
		serviceName: "telegram",
	}
	// Check if default chat ID is set
	if config.ChatID != "" {
		parsedChatID, err := strconv.ParseInt(config.ChatID, 10, 64)
		if err != nil {
			namedLogger.Error(context.Background(), "Invalid ChatID in Telegram config", map[string]any{"chat_id": config.ChatID, "error": err})
		} else {
			adapter.chatID = parsedChatID
			namedLogger.Info(context.Background(), "Telegram adapter configured with default chat ID", map[string]any{"chat_id": parsedChatID})
		}
	}

	namedLogger.Info(context.Background(), "Telegram notify adapter (library) initialized", map[string]any{"bot_username": bot.Self.UserName})
	return adapter, nil
}

// Send a message via Telegram using the library.
func (a *TelegramAdapter) Send(ctx context.Context, msg dto.Content) error {
	// 1) Fallback parse-mode
	if msg.ParseMode == "" {
		msg.ParseMode = tgbotapi.ModeHTML
	}

	// 2) Pick icon
	levelIcon := map[dto.Level]string{
		dto.Debug:   "🔍",
		dto.Info:    "ℹ️",
		dto.Warning: "⚠️",
		dto.Error:   "❌",
	}[msg.Level]
	if levelIcon == "" {
		levelIcon = "📢"
	}

	// 3) Build safe HTML (no <span>, no style)
	subject := html.EscapeString(msg.Subject)
	body := html.EscapeString(msg.Message)

	var text string
	if subject != "" {
		text = fmt.Sprintf("<b>%s %s</b>\n\n<pre>%s</pre>\n\n<i>Level: %s</i>",
			levelIcon, subject, body, msg.Level)
	} else {
		text = fmt.Sprintf("<b>%s Notification</b>\n\n<pre>%s</pre>\n\n<i>Level: %s</i>",
			levelIcon, body, msg.Level)
	}

	m := tgbotapi.NewMessage(a.chatID, text)
	m.ParseMode = msg.ParseMode

	if _, err := a.bot.Send(m); err != nil {
		a.logger.Error(ctx, "telegram send failed", map[string]any{"error": err})
		return err
	}
	a.logger.Debug(ctx, "telegram send ok", nil)
	return nil
}
