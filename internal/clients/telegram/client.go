package telegram

// Client for interaction with telegram API
type Client struct {
	api BotAPI
}

// NewTelegramClient creates a new telegram client for receving and sending messages
func NewTelegramClient(api BotAPI) *Client {
	return &Client{api: api}
}
