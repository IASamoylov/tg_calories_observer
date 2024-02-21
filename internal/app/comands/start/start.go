package start

import (
	"context"
	"fmt"
	"math"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/IASamoylov/tg_calories_observer/internal/utils/types"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/logger"

	"github.com/samber/lo"

	"github.com/IASamoylov/tg_calories_observer/internal/domain/entity/dto"
)

const (
	name = "start"
	text = "👋 Здравствуй, @%s! Я - твой персональный калорийный агент 🍩🍩5 калл. " +
		"Моя миссия - отслеживать и контролировать твое суточное потребление калорий. " +
		"Я ищу 'злодеев' вроде неполезных продуктов 🍟🍫, но без твоей помощи я не справлюсь. " +
		"Давай вместе вести дневник раследований, чтобы выявить все 'пострадавшие' продукты 🍎. \n\n" +
		"Чтобы узнать подробнее что я умею, необходимо ввести команду /%s."
	help = "\n\nДля твоего удобства я подготовил меню быстрого доступа 👇"

	description = "Показать привестивие"
)

// Handler обработчик команды /start
type Handler struct {
	helpCmd        helpCommand
	keyboard       []keyboardButton
	telegram       types.Telegram
	profilePhotoID *string
}

// Execute исполняет команду /start
func (handler *Handler) Execute(_ context.Context, user dto.User, _ string) (tgbotapi.Chattable, error) {
	keyboard := handler.getReplyMarkup()

	if fileID := handler.getFileID(); fileID != nil {
		msg := tgbotapi.NewPhoto(user.TelegramID(), tgbotapi.FileID(*fileID))
		msg.Caption = fmt.Sprintf(text, user.UserName(), handler.helpCmd.Alias())
		if len(handler.keyboard) > 0 {
			msg.ReplyMarkup = keyboard
			msg.Caption += help
		}

		return msg, nil
	}

	msg := tgbotapi.NewMessage(user.TelegramID(), fmt.Sprintf(text, user.UserName(), handler.helpCmd.Alias()))
	if len(handler.keyboard) > 0 {
		msg.ReplyMarkup = keyboard
		msg.Text += help
	}

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

// WithKeyboardButton добавляет клавиатуру, которая упростит использование приложения
func (handler *Handler) WithKeyboardButton(buttons ...keyboardButton) {
	handler.keyboard = append(handler.keyboard, buttons...)
}

func (handler *Handler) getFileID() *string {
	if handler.profilePhotoID != nil {
		return handler.profilePhotoID
	}

	me, err := handler.telegram.GetMe()
	if err != nil {
		logger.Error(fmt.Sprintf("Ошибка получения профиля бота: %s", err), "command", name)

		return nil
	}

	files, err := handler.telegram.GetUserProfilePhotos(tgbotapi.UserProfilePhotosConfig{UserID: me.ID, Limit: 1})
	if err != nil {
		logger.Error(fmt.Sprintf("Ошибка при получение фото профиля бота: %s", err), "command", name)

		return nil
	}

	if len(files.Photos) == 0 || len(files.Photos[0]) == 0 {
		logger.Error("Фотографии в профиле не найдены", "command", name)

		return nil
	}

	photo := files.Photos[0][0]

	handler.profilePhotoID = &photo.FileID

	return handler.profilePhotoID
}

func (handler *Handler) getReplyMarkup() tgbotapi.ReplyKeyboardMarkup {
	if len(handler.keyboard) > 0 {
		rows := lo.Chunk(handler.keyboard, int(math.Ceil(float64(len(handler.keyboard))/2)))
		keyboard := tgbotapi.NewReplyKeyboard(lo.Map(rows, func(row []keyboardButton, _ int) []tgbotapi.KeyboardButton {
			return tgbotapi.NewKeyboardButtonRow(lo.Map(row, func(item keyboardButton, _ int) tgbotapi.KeyboardButton {
				return tgbotapi.NewKeyboardButton(item.Text())
			})...)
		})...)
		keyboard.OneTimeKeyboard = true

		return keyboard
	}

	return tgbotapi.ReplyKeyboardMarkup{}
}

// NewHandler ctor
func NewHandler(helpCmd helpCommand, api types.Telegram) *Handler {
	return &Handler{helpCmd: helpCmd, telegram: api}
}
