package help

import (
	"context"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/stretchr/testify/assert"

	"github.com/IASamoylov/tg_calories_observer/internal/domain/entity/dto"

	"go.uber.org/mock/gomock"
)

func TestExecute(t *testing.T) {
	t.Parallel()

	t.Run("команда возвращает подсказку пользователю со списком возможных команд", func(t *testing.T) {
		t.Parallel()

		mockCommand := NewMockсommand(gomock.NewController(t))

		mockCommand.EXPECT().Alias().Return("cmd_1")
		mockCommand.EXPECT().Description().Return("описание команды 1")
		mockCommand.EXPECT().Alias().Return("cmd_2")
		mockCommand.EXPECT().Description().Return("описание команды 2")

		handler := NewHandler(mockCommand, mockCommand)

		user := dto.NewUser(100500, "", "", "", "")

		msg, err := handler.Execute(context.Background(), user, "")
		assert.NoError(t, err)

		expectedMessage := "Вот что я умею:\n" +
			"- /cmd_1 - описание команды 1 \n" +
			"- /cmd_2 - описание команды 2 \n\n" +
			"Если у вас есть вопросы или нужна помощь, не стесняйтесь обращаться!"
		assert.EqualValues(t, tgbotapi.NewMessage(100500, expectedMessage), msg)
	})

	t.Run("если подсказки не проинцилизированы команда не отработает", func(t *testing.T) {
		t.Parallel()

		handler := NewHandler()
		user := dto.NewUser(100500, "", "", "", "")

		msg, err := handler.Execute(context.Background(), user, "")
		assert.NoError(t, err)
		expectedMessage := "Вот что я умею:\n\n" +
			"Если у вас есть вопросы или нужна помощь, не стесняйтесь обращаться!"
		assert.EqualValues(t, tgbotapi.NewMessage(100500, expectedMessage), msg)
	})
}

func TestAlias(t *testing.T) {
	t.Parallel()

	t.Run("возврщается название команды", func(t *testing.T) {
		handler := NewHandler()

		assert.Equal(t, "help", handler.Alias())
	})
}

func TestDescription(t *testing.T) {
	t.Parallel()

	t.Run("возврщается описание команды", func(t *testing.T) {
		handler := NewHandler()

		assert.Equal(t, "Попросить помощи", handler.Description())
	})
}
