package telegram

import (
	"github.com/IASamoylov/tg_calories_observer/internal/pkg/logger"
	"github.com/IASamoylov/tg_calories_observer/internal/utils/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Client клиент для взаимодействия с API телеграмма
type Client struct {
	telegram types.Telegram
}

// NewClient ctor
func NewClient(telegram types.Telegram) *Client {
	return &Client{telegram: telegram}
}

// Send отправляет сообщение пользователю
func (client *Client) Send(c tgbotapi.Chattable) {
	resp, err := client.telegram.Send(c)
	if err != nil {
		logger.Error("При выполненеи 'Send' произошла ошибка", "err", err, "req", c, "resp", resp)
	}
}

// Request выполняет запрос к АПИ
func (client *Client) Request(c tgbotapi.Chattable) {
	resp, err := client.telegram.Request(c)
	if err != nil {
		logger.Error("При выполненеи 'Request' произошла ошибка", "err", err, "req", c, "resp", resp)
	}
}

// SendErr отправляет сообщение с ошибкой пользователю
func (client *Client) SendErr(receiver int64, err error) {
	msg := tgbotapi.NewMessage(receiver, err.Error())
	msg.ParseMode = tgbotapi.ModeMarkdownV2

	client.Send(msg)
}

// InitMenu инцилизирует меню в телеграм
func (client *Client) InitMenu(commands []tgbotapi.BotCommand) {
	client.Request(tgbotapi.NewDeleteMyCommands())
	client.Request(tgbotapi.NewSetMyCommands(commands...))
}
