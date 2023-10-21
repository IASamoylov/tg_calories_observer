package telegram

import (
	"github.com/IASamoylov/tg_calories_observer/internal/pkg/types"
)

// Controller a set of handles for receiving messages from telegram view webhook
type Controller struct {
	bot types.TelegramBotAPI
}

// NewController ctor
func NewController(bot types.TelegramBotAPI) Controller {
	return Controller{
		bot: bot,
	}
}
