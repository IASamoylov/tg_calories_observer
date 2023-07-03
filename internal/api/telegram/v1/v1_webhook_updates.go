package telegram

import (
	"fmt"
	"io"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/IASamoylov/tg_calories_observer/internal/config/debug"
)

// V1WebhookUpdates receiving a message from telegram using a web hook
func (ctr Controller) V1WebhookUpdates(writer http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		return
	}

	update := <-ctr.bot.ListenForWebhookRespReqFormat(writer, req)
	log.Println(update)
	go func(update tgbotapi.Update) {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msgText := fmt.Sprintf(`[%s] Привет %s (@%s), ты прислал мне сообщение "%s", но сделал это без уважения`,
				debug.Version, update.Message.From.FirstName, update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
			_, _ = ctr.bot.Send(msg)
			photo := requestPhoto()
			if len(photo) != 0 {
				msg := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileBytes{
					Name:  "The Godfather",
					Bytes: photo,
				})
				_, _ = ctr.bot.Send(msg)

			}
		}
	}(update)
}

func requestPhoto() []byte {
	resp, err := http.Get("https://static1.colliderimages.com/wordpress/wp-content/" +
		"uploads/2022/11/The-Godfather.jpg?q=50&fit=contain&w=1140&h=&dpr=1.5")
	if err == nil && resp.StatusCode == http.StatusOK {
		// nolint
		defer resp.Body.Close()

		photo, _ := io.ReadAll(resp.Body)

		return photo
	}

	return []byte{}
}
