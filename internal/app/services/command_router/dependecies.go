package commandrouter

import (
	"context"

	"github.com/IASamoylov/tg_calories_observer/internal/domain/entity/dto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//go:generate mockgen -source=dependecies.go -destination=dependecies_mocks.go -package=command_router . userStorage,telegram,commandRouterHandler

type userStorage interface {
	Upsert(ctx context.Context, user dto.User) error
}

type telegram interface {
	Send(c tgbotapi.Chattable)
	SendErr(receiver int64, err error)
	InitMenu(commands []tgbotapi.BotCommand)
}

type commandRouterHandler interface {
	Execute(ctx context.Context, sender dto.User, args ...string) (tgbotapi.Chattable, error)
	Alias() string
	Description() string
}
