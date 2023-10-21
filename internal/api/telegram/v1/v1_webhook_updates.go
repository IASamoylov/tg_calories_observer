package telegram

import (
	"net/http"
)

// V1WebhookUpdates receiving a message from telegram using a web hook
func (ctr Controller) V1WebhookUpdates(writer http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		return
	}

	//update := <-ctr.bot.ListenForWebhookRespReqFormat(writer, req)

	//if update.Message == nil || { // If we got a message
	//	return
	//}
	//
	//if !update.Message.IsCommand() { // If we got a message
	//	return
	//}
	//
	//if update.CallbackQuery != nil {
	//
	//}
}
