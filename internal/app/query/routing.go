package query

import (
	"context"
	"fmt"
	"strings"

	"github.com/IASamoylov/tg_calories_observer/internal/domain"
)

// Routing a router of commands that must somehow modify, create or delete domain entities
type Routing struct {
	route map[string]Handler
}

func NewRouting() *Routing {
	return &Routing{
		route: make(map[string]Handler),
	}
}

type Handler interface {
	Execute(ctx context.Context, sender domain.User, args []any) (
		recipient domain.User,
		msg string,
		err error,
	)
	Parse(message string) (args []any)
	Alias() string
}

func (routing *Routing) Add(handler Handler) *Routing {
	routing.route[handler.Alias()] = handler
	routing.route[fmt.Sprintf("/%s", handler.Alias())] = handler

	return routing
}

func (routing *Routing) Execute(ctx context.Context, user domain.User, message string) (
	recipient domain.User,
	msg string,
	err error,
) {
	handler, _ := routing.getHandler(message)

	return handler.Execute(ctx, user, handler.Parse(message))
}

func (routing *Routing) IsQuery(message string) bool {
	_, ok := routing.getHandler(message)

	return ok
}

func (routing *Routing) getHandler(message string) (Handler, bool) {
	handler, ok := routing.route[message]
	if ok {
		return handler, ok
	}

	for alias, handler := range routing.route {
		if ok = strings.HasPrefix(message, alias); ok {
			return handler, ok
		}
	}

	return nil, false
}
