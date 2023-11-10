package product

import (
	"context"
	"testing"

	"github.com/samber/lo"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/IASamoylov/tg_calories_observer/internal/domain/entity/dto"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

var user = dto.NewUser(100500, "test", "", "", "")

func TestAlias(t *testing.T) {
	t.Parallel()

	t.Run("возвращает текст, который должен отобразиться у пользователя на кнопке", func(t *testing.T) {
		t.Parallel()

		button := NewButton(nil)

		assert.Equal(t, "🍇Продукты", button.Text())
	})
}

func TestExecute(t *testing.T) {
	t.Parallel()

	t.Run("возвращает сообщение и список команд, которые пользователь может выполнить", func(t *testing.T) {
		t.Parallel()

		btn := NewMockinlineButton(gomock.NewController(t))
		btn.EXPECT().Text().Return("Добавить")
		btn.EXPECT().Callback().Return("/add")
		btn.EXPECT().Text().Return("Удалить")
		btn.EXPECT().Callback().Return("/remove")
		btn.EXPECT().Text().Return("Изменить")
		btn.EXPECT().Callback().Return("/edit")

		button := NewButton(btn, btn, btn)

		msg, err := button.Execute(context.Background(), user)

		expectedMessage := tgbotapi.NewMessage(user.TelegramID(), "Какое действие необходимо выполнить?")
		expectedMessage.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{
			InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
				{
					{Text: "Добавить", CallbackData: lo.ToPtr("/add")},
					{Text: "Удалить", CallbackData: lo.ToPtr("/remove")},
				},
				{
					{Text: "Изменить", CallbackData: lo.ToPtr("/edit")},
				},
			},
		}

		assert.NoError(t, err)
		assert.Equal(t, expectedMessage, msg)
	})
}
