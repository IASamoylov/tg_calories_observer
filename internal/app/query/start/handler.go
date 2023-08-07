package start

import (
	"context"
	"fmt"
	"strings"

	"github.com/IASamoylov/tg_calories_observer/internal/domain"
)

const alias = "start"
const description = "Попросить помощи"

type Handler struct {
}

// Execute executes command
func (handler *Handler) Execute(_ context.Context, user domain.User, _ []any) (domain.User, string, error) {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Привет @%s, я бот помощник для контроля суточной "+
		"нормы калорий и БЖУ, вот что я умею:", user.UserName()))
	sb.WriteString(fmt.Sprintf(" - %s", handler.String()))

	return user, sb.String(), nil
}

func (handler *Handler) Parse(_ string) (_ []any) {
	return nil
}

func (handler *Handler) String() string {
	return fmt.Sprintf("%s - %s", alias, description)
}

// Alias command name
func (handler *Handler) Alias() string {
	return alias
}

// Description describes the behavior of the command
func (handler *Handler) Description() string {
	return description
}

func NewQueryHandler() *Handler {
	return &Handler{}
}
