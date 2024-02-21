package product

import (
	"context"
	"math"

	dto "github.com/IASamoylov/tg_calories_observer/internal/domain/entity/dto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/samber/lo"
)

const (
	name = "🍇Продукты"
	text = "Какое действие необходимо выполнить?"
)

// Handler обработчик кнопки
type Handler struct {
	keyboard []inlineButton
}

// NewButton ctor
func NewButton(keyboard ...inlineButton) *Handler {
	return &Handler{keyboard: keyboard}
}

// Execute выполняет команду, результатом которой будет ответ пользователю
func (handler *Handler) Execute(_ context.Context, sender dto.User) (tgbotapi.Chattable, error) {
	msg := tgbotapi.NewMessage(sender.TelegramID(), text)
	rows := lo.Chunk(handler.keyboard, int(math.Ceil(float64(len(handler.keyboard))/2)))

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		lo.Map(rows, func(row []inlineButton, _ int) []tgbotapi.InlineKeyboardButton {
			return tgbotapi.NewInlineKeyboardRow(
				lo.Map(row, func(btn inlineButton, _ int) tgbotapi.InlineKeyboardButton {
					return tgbotapi.NewInlineKeyboardButtonData(btn.Text(), btn.Callback())
				})...,
			)
		})...,
	)

	return msg, nil
}

// Text возврщает текст кнопки
func (handler *Handler) Text() string {
	return name
}

// keyboard
// [Продукты] -> Inline keyboard
//               [Добавить] callback data {/add_product} ?state_machine?
//               [Удалить] callback data {/remove_product}
//               [Посмотреть] callback data {/get_product}
//               [Изменить] callback data {/edit_product} -> Inline keyboard ?state_machine?
//               											 [калл] callback data {/edit_product call}
//               											 [название] callback data {/edit_product name}
//               											 [б] callback data {/edit_product proteins}
//               											 [ж] callback data {/edit_product fats}
//               											 [у] callback data {/edit_product carbohydrates}
