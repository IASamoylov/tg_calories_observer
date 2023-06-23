package telegram

// Client for interaction with telegram API
type Client struct {
	api TelegramBotAPI
}

// NewTelegramClient creates a new telegram client for receving and sending messages
func NewTelegramClient(api TelegramBotAPI) *Client {
	return &Client{api: api}
}
