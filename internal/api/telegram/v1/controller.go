package telegram

import (
	"github.com/IASamoylov/tg_calories_observer/internal/app/services/message_routing"
	"github.com/IASamoylov/tg_calories_observer/internal/pkg/types"
)

// Controller a set of handles for receving messages from telegram view webhook
type Controller struct {
	bot            types.TelegramBotAPI
	messageRouting message_routing.Service
}

// NewController ctor
func NewController(bot types.TelegramBotAPI, messageRouting message_routing.Service) *Controller {
	return &Controller{
		bot:            bot,
		messageRouting: messageRouting,
	}
}
