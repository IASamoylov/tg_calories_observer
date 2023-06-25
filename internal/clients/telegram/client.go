package telegram

import "github.com/IASamoylov/tg_calories_observer/internal/pkg/types"

// Client for interaction with telegram API
type Client struct {
	bot types.TelegramBotAPI
}

// NewTelegramClient creates a new telegram client for receving and sending messages
func NewTelegramClient(api types.TelegramBotAPI) *Client {
	client := &Client{bot: api}

	return client
}
