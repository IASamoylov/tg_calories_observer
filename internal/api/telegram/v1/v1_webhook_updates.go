package telegram

import (
	"net/http"

	"github.com/IASamoylov/tg_calories_observer/internal/domain"
)

// V1WebhookUpdates receiving a message from telegram using a web hook
func (ctr Controller) V1WebhookUpdates(writer http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		return
	}

	update := <-ctr.bot.ListenForWebhookRespReqFormat(writer, req)
	if update.Message != nil { // If we got a message
		from := update.Message.From
		user := domain.NewDefaultUser(from.ID, from.UserName, from.FirstName, from.LastName, from.LanguageCode)

		ctr.messageRouting.Handle(req.Context(), user, update.Message.Text)
	}
}
