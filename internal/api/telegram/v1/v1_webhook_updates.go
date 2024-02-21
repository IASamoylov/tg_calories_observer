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

	update := <-ctr.decoder.ListenForWebhookRespReqFormat(writer, req)
	if update.Message == nil && update.CallbackQuery == nil { // If we got a message
		return
	}

	if update.CallbackQuery != nil {
		//ctr.commandRouter.Execute(req.Context(), getUser(update.CallbackQuery.From), update.CallbackQuery.Data)

		update.Message = &tgbotapi.Message{
			From: update.CallbackQuery.From,
			Text: update.CallbackQuery.Data,
		}
	}

	if update.Message.IsCommand() {
		ctr.commandRouter.Execute(
			req.Context(),
			dto.NewUser(
				update.Message.From.ID,
				update.Message.From.UserName,
				update.Message.From.FirstName,
				update.Message.From.LastName,
				update.Message.From.LanguageCode,
			),
			update.Message.Command(),
			update.Message.CommandArguments(),
		)

		return
	}

	//ctr.keyboardRouter.Execute(req.Context(), getUser(update.Message.From), update.Message.Text)
}
