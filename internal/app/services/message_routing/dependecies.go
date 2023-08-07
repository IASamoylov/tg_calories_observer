package message_routing

import (
	"context"

	"github.com/IASamoylov/tg_calories_observer/internal/domain"
)

type RequestRouting interface {
	IsQuery(message string) bool
	Execute(ctx context.Context, user domain.User, message string) (
		recipient domain.User,
		msg string,
		err error,
	)
}

type CommandRouting interface {
	IsCommand(message string) bool
	Execute(ctx context.Context, command any) any
}

type MessengerClient interface {
	Send(recipient domain.User, text string) error
}

type UserGetter interface {
	UpsertAndGet(ctx context.Context, user domain.User) (domain.User, error)
}
