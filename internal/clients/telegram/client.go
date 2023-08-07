package telegram

import (
	"github.com/samber/lo"

	"github.com/IASamoylov/tg_calories_observer/internal/domain"
	"github.com/IASamoylov/tg_calories_observer/internal/pkg/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Client for interaction with telegram API
type Client struct {
	bot types.TelegramBotAPI
}

// Send sends message to chat or user
func (client Client) Send(user domain.User, text string) error {
	msg := tgbotapi.NewMessage(user.TelegramID(), text)
	_, err := client.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

// InitMenu inits help menu for telegram bot
func (client Client) InitMenu(buttons ...Button) error {
	_, err := client.bot.Send(tgbotapi.SetMyCommandsConfig{
		Commands: lo.Map(buttons, func(cmd Button, _ int) tgbotapi.BotCommand {
			return tgbotapi.BotCommand{Command: cmd.Alias(), Description: cmd.Description()}
		}),
	})

	return err
}

// NewTelegramClient creates a new telegram client for receving and sending messages
func NewTelegramClient(api types.TelegramBotAPI) *Client {
	client := &Client{bot: api}

	return client
}
