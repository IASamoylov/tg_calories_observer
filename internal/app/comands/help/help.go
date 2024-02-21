package help

import (
	"context"
	"fmt"
	"strings"

	"github.com/IASamoylov/tg_calories_observer/internal/domain/entity/dto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	name = "help"
	text = "Вот что я умею:\n" +
		"%s\n" +
		"Если у вас есть вопросы или нужна помощь, не стесняйтесь обращаться!"
	description = "Попросить помощи"
)

// Handler обработчик команды /help
type Handler struct {
	commands []сommand
}

// Execute исполняет команду /help
func (handler *Handler) Execute(_ context.Context, user dto.User, _ string) (tgbotapi.Chattable, error) {
	var builder strings.Builder
	for _, cmd := range handler.commands {
		builder.WriteString(fmt.Sprintf("- /%s - %s \n", cmd.Alias(), cmd.Description()))
	}

	msg := tgbotapi.NewMessage(user.TelegramID(), fmt.Sprintf(text, builder.String()))

	return msg, nil
}

// Alias возвращает название команды
func (handler *Handler) Alias() string {
	return name
}

// Description возвращает описание команды
func (handler *Handler) Description() string {
	return description
}

// NewHandler ctor
func NewHandler(commands ...сommand) *Handler {
	return &Handler{commands: commands}
}
