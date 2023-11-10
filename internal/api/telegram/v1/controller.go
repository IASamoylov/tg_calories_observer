package telegram

import (
	commandrouter "github.com/IASamoylov/tg_calories_observer/internal/app/services/command_router"
	"github.com/IASamoylov/tg_calories_observer/internal/utils/types"
)

// Controller a set of handles for receiving messages from telegram view webhook
type Controller struct {
	bot           types.Telegram
	commandRouter *commandrouter.CommandRouter
	//keyboardRouter *routers.KeyboardRouter
}

// NewController ctor
func NewController(
	bot types.Telegram,
	commandRouter *commandrouter.CommandRouter,
	// keyboardRouter *routers.KeyboardRouter,
) Controller {
	return Controller{
		bot:           bot,
		commandRouter: commandRouter,
		//keyboardRouter: keyboardRouter,
	}
}
