package telegram

import (
	"github.com/IASamoylov/tg_calories_observer/internal/pkg/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Client for interaction with telegram API
type Client struct {
	bot types.TelegramBotAPI
}

// SendMsg sends message to chat or user
func (client Client) SendMsg(recipient int64, text string) error {
	msg := tgbotapi.NewMessage(recipient, text)
	_, err := client.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

// NewTelegramClient creates a new telegram client for receving and sending messages
func NewTelegramClient(api types.TelegramBotAPI) Client {
	client := Client{bot: api}

	return client
}
