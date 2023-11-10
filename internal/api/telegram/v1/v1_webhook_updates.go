package telegram

import (
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/IASamoylov/tg_calories_observer/internal/domain/entity/dto"
)

// V1WebhookUpdates receiving a message from telegram using a web hook
func (ctr Controller) V1WebhookUpdates(writer http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		return
	}

	update := <-ctr.bot.ListenForWebhookRespReqFormat(writer, req)

	if update.Message == nil && update.CallbackQuery == nil { // If we got a message
		return
	}

	getUser := func(from *tgbotapi.User) dto.User {
		return dto.NewUser(
			from.ID,
			from.UserName,
			from.FirstName,
			from.LastName,
			from.LanguageCode,
		)
	}

	if update.CallbackQuery != nil {
		ctr.commandRouter.Execute(req.Context(), getUser(update.CallbackQuery.From), update.CallbackQuery.Data)

		return
	}

	if update.Message.IsCommand() {
		ctr.commandRouter.Execute(req.Context(), getUser(update.Message.From), update.Message.Text)

		return
	}

	//ctr.keyboardRouter.Execute(req.Context(), getUser(update.Message.From), update.Message.Text)
}
