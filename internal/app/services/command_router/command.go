package commandrouter

import (
	"context"
	"fmt"
	"strings"

	"github.com/IASamoylov/tg_calories_observer/internal/utils/types"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/IASamoylov/tg_calories_observer/internal/domain/entity/dto"
	"github.com/samber/lo"
)

// CommandRouter маршрутизатор комманд из чата. Запускает обоработчики на основание алиаса команды
type CommandRouter struct {
	route       map[string]commandRouterHandler
	client      telegram
	userStorage userStorage
	handlers    []commandRouterHandler
	helpCommand commandRouterHandler
}

// NewCommandRouter ctor
func NewCommandRouter(
	client telegram,
	userStorage userStorage,
	helpCommand commandRouterHandler,
	handlers ...commandRouterHandler,
) *CommandRouter {
	return &CommandRouter{
		handlers:    handlers,
		helpCommand: helpCommand,
		route: lo.SliceToMap(handlers, func(handler commandRouterHandler) (string, commandRouterHandler) {
			return handler.Alias(), handler
		}),
		userStorage: userStorage,
		client:      client,
	}
}

// Execute маршрутизирует команду
func (router *CommandRouter) Execute(ctx context.Context, user dto.User, message string) {
	defer router.recover(user)

	err := router.userStorage.Upsert(ctx, user)
	if err != nil {
		logger.Error("Не удалось обновить данные пользователя", "err", err)
		router.client.SendErr(user.TelegramID(), types.ErrCommon)

		return
	}

	parts := strings.Split(message, " ")
	if len(parts) == 0 || len(message) == 0 {
		return
	}

	if handler, ok := router.route[parts[:1][0]]; ok {
		msg, err := handler.Execute(ctx, user, parts[1:]...)
		if err != nil {
			logger.Error("Возникла ошибки в момент обработки команды", "err", err, "command", parts[:1][0])
			router.client.SendErr(user.TelegramID(), types.ErrCommon)

			return
		}
		router.client.Send(msg)

		return
	}

	err = fmt.Errorf("@%s к сожалению, ваш запрос '%s' не распознан. Прошу прощения за неудобства. "+
		"Продолжайте следить за инструкциями %s и оставайтесь "+
		"на связи для дальнейших указаний.", user.UserName(), parts[:1][0], router.helpCommand.Alias())

	router.client.SendErr(user.TelegramID(), err)
}

// InitMenu инцилизирует меню в телеграм
func (router *CommandRouter) InitMenu() {
	commands := lo.Map(router.handlers, func(cmd commandRouterHandler, _ int) tgbotapi.BotCommand {
		return tgbotapi.BotCommand{Command: cmd.Alias(), Description: cmd.Description()}
	})

	router.client.InitMenu(commands)
}

func (router *CommandRouter) recover(user dto.User) {
	if detail := recover(); detail != nil {
		logger.Errorf("Произошла непридведенная ошибка: %s", detail)
		router.client.SendErr(user.TelegramID(), types.ErrCommon)
	}
}
