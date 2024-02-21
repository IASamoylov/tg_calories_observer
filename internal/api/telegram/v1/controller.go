package telegram

import (
	commandrouter "github.com/IASamoylov/tg_calories_observer/internal/app/services/command_router"
	"github.com/IASamoylov/tg_calories_observer/internal/utils/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Controller a set of handles for receiving messages from telegram view webhook
type Controller struct {
	decoder       types.Telegram
	commandRouter *commandrouter.CommandRouter
	//keyboardRouter *routers.KeyboardRouter
}

// NewController ctor
func NewController(
	commandRouter *commandrouter.CommandRouter,
	// keyboardRouter *routers.KeyboardRouter,
) Controller {
	return Controller{
		decoder: &tgbotapi.BotAPI{
			Buffer: 100,
		},
		commandRouter: commandRouter,
		//keyboardRouter: keyboardRouter,
	}
}
