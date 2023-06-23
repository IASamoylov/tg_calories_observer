package telegram

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"

	telegrambotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/IASamoylov/tg_calories_observer/internal/config/debug"
)

// Client for interaction with telegram API
type Client struct {
	api BotAPI
}

var url = "https://static1.colliderimages.com/wordpress/wp-content/uploads/2022/11/" +
	"The-Godfather.jpg?q=50&fit=contain&w=1140&h=&dpr=1.5"

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

		resp, err := http.Get(url)
		defer resp.Body.Close()

		var reader io.Reader

		if err == nil && resp.StatusCode == http.StatusOK {
			reader = bufio.NewReader(resp.Body)
		}

		for update := range updates {
			if update.Message != nil { // If we got a message
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

				msgText := fmt.Sprintf(`[%s] Привет %s (@%s), ты прислал мне сообщение "%s", но сделал это без уважения`,
					debug.Version, update.Message.From.FirstName, update.Message.From.UserName, update.Message.Text)

				msg := telegrambotapi.NewMessage(update.Message.Chat.ID, msgText)
				_, _ = api.Send(msg)
				if reader != nil {
					msg := telegrambotapi.NewPhoto(update.Message.Chat.ID, telegrambotapi.FileReader{
						Name:   "The Godfather",
						Reader: reader,
					})
					_, _ = api.Send(msg)

				}
			}
		}
	}()

	return &Client{api: api}
}
