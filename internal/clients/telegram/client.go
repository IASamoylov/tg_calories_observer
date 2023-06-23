package telegram

import (
	"log"

	telegrambotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Client for interaction with telegram API
type Client struct {
	api BotAPI
}

// NewTelegramClient creates a new telegram client for receving and sending messages
func NewTelegramClient(api BotAPI) *Client {

	go func() {

		self, err := api.GetMe()
		if err != nil {
			log.Println(err)

			return
		}
		log.Printf("Authorized on account %s", self.UserName)

		u := telegrambotapi.NewUpdate(0)
		u.Timeout = 60

		updates := api.GetUpdatesChan(u)

		for update := range updates {
			if update.Message != nil { // If we got a message
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

				msg := telegrambotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.ReplyToMessageID = update.Message.MessageID

				_, _ = api.Send(msg)
			}
		}
	}()

	return &Client{api: api}
}
