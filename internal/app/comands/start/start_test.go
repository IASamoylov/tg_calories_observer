package start

import (
	"context"
	"fmt"
	"testing"

	mockstelegram "github.com/IASamoylov/tg_calories_observer/internal/utils/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/IASamoylov/tg_calories_observer/internal/domain/entity/dto"

	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
)

const expectedText = "👋 Здравствуй, @test! Я - твой персональный калорийный агент 🍩🍩5 калл. " +
	"Моя миссия - отслеживать и контролировать твое суточное потребление калорий. " +
	"Я ищу 'злодеев' вроде неполезных продуктов 🍟🍫, но без твоей помощи я не справлюсь. " +
	"Давай вместе вести дневник раследований, чтобы выявить все 'пострадавшие' продукты 🍎. \n\n" +
	"Чтобы узнать подробнее что я умею, необходимо ввести команду /help."

var user = dto.NewUser(100500, "test", "", "", "")
var me = tgbotapi.User{ID: 35125}

func beforeTest(t *testing.T) (*Handler, *MockkeyboardButton, *mockstelegram.MockTelegram) {
	mockHelpCommand := NewMockhelpCommand(gomock.NewController(t))
	mockHelpCommand.EXPECT().Alias().Return("help").AnyTimes()

	mockKeyboardButton := NewMockkeyboardButton(gomock.NewController(t))
	mockTelegram := mockstelegram.NewMockTelegram(gomock.NewController(t))

	return NewHandler(mockHelpCommand, mockTelegram), mockKeyboardButton, mockTelegram
}

func TestExecute(t *testing.T) {
	t.Parallel()

	t.Run("команда возвращает привествие пользователю с фото профиля бота", func(t *testing.T) {
		t.Parallel()

		handler, _, mockTelegram := beforeTest(t)
		fileID := "cb196684-d6ec-4fb1-a5bf-c87e39193514"
		expectedPhoto := tgbotapi.NewPhoto(user.TelegramID(), tgbotapi.FileID(fileID))
		expectedPhoto.Caption = expectedText

		mockTelegram.EXPECT().GetMe().Return(me, nil)
		mockTelegram.EXPECT().
			GetUserProfilePhotos(tgbotapi.UserProfilePhotosConfig{UserID: me.ID, Limit: 1}).
			Return(tgbotapi.UserProfilePhotos{Photos: [][]tgbotapi.PhotoSize{{{FileID: fileID}}}}, nil).
			Times(1)

		msg, err := handler.Execute(context.Background(), user, "")
		assert.NoError(t, err)
		assert.Equal(t, expectedPhoto, msg)

		t.Run("повторный вызов команды не будет запрашивать фото профиля", func(t *testing.T) {
			msg, err = handler.Execute(context.Background(), user, "")
			assert.NoError(t, err)

			assert.Equal(t, expectedPhoto, msg)
		})
	})

	t.Run("если фото профиля не удалось получить, то команда возвращает только приветствие", func(t *testing.T) {
		t.Parallel()

		t.Run("не удалось получить профиль бота", func(t *testing.T) {
			t.Parallel()

			handler, _, mockTelegram := beforeTest(t)
			expectedMessage := tgbotapi.NewMessage(user.TelegramID(), expectedText)

			mockTelegram.EXPECT().GetMe().Return(tgbotapi.User{}, fmt.Errorf("GetMe error"))
			mockTelegram.EXPECT().GetUserProfilePhotos(gomock.Any()).Times(0)

			msg, err := handler.Execute(context.Background(), user, "")
			assert.NoError(t, err)
			assert.Equal(t, expectedMessage, msg)
		})

		t.Run("не удалось получить фото профиля бота", func(t *testing.T) {
			t.Parallel()

			handler, _, mockTelegram := beforeTest(t)
			expectedMessage := tgbotapi.NewMessage(user.TelegramID(), expectedText)

			mockTelegram.EXPECT().GetMe().Return(me, nil)
			mockTelegram.EXPECT().
				GetUserProfilePhotos(tgbotapi.UserProfilePhotosConfig{UserID: me.ID, Limit: 1}).
				Return(tgbotapi.UserProfilePhotos{}, fmt.Errorf("GetUserProfilePhotos error"))

			msg, err := handler.Execute(context.Background(), user, "")

			assert.NoError(t, err)
			assert.Equal(t, expectedMessage, msg)
		})

		t.Run("у профили бота отсутсвуют фото", func(t *testing.T) {
			t.Parallel()

			tcs := []struct {
				name  string
				photo [][]tgbotapi.PhotoSize
			}{
				{name: "nil", photo: nil},
				{name: "empty", photo: [][]tgbotapi.PhotoSize{}},
				{name: "отсутвует фото любого размера", photo: [][]tgbotapi.PhotoSize{{}}},
			}

			for _, tc := range tcs {
				tc := tc

				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					handler, _, mockTelegram := beforeTest(t)
					expectedMessage := tgbotapi.NewMessage(user.TelegramID(), expectedText)

					mockTelegram.EXPECT().GetMe().Return(me, nil)
					mockTelegram.EXPECT().
						GetUserProfilePhotos(tgbotapi.UserProfilePhotosConfig{UserID: me.ID, Limit: 1}).
						Return(tgbotapi.UserProfilePhotos{Photos: tc.photo}, nil)

					msg, err := handler.Execute(context.Background(), user, "")

					assert.NoError(t, err)
					assert.Equal(t, expectedMessage, msg)
				})
			}
		})
	})
}

func TestWithKeyboardButton(t *testing.T) {
	t.Parallel()

	expectedText := expectedText + help
	expectedProductBtn := tgbotapi.KeyboardButton{Text: "Продукты"}
	expectedReportBrn := tgbotapi.KeyboardButton{Text: "Отчеты"}
	expectedSettingBtn := tgbotapi.KeyboardButton{Text: "Настройки"}
	expectedRow1 := []tgbotapi.KeyboardButton{expectedProductBtn, expectedReportBrn}
	expectedRow2 := []tgbotapi.KeyboardButton{expectedSettingBtn}
	expectedKeyboard := tgbotapi.ReplyKeyboardMarkup{
		OneTimeKeyboard: true,
		ResizeKeyboard:  true,
		Keyboard:        [][]tgbotapi.KeyboardButton{expectedRow1, expectedRow2},
	}

	t.Run("если настроена клавиаутура она будет показана пользователю вместе с сообщением", func(t *testing.T) {
		t.Parallel()

		t.Run("только приветсвие при отсутвие фотографии профиля бота", func(t *testing.T) {
			t.Parallel()
			handler, mockKeyboardButton, mockTelegram := beforeTest(t)

			expectedMessage := tgbotapi.NewMessage(user.TelegramID(), expectedText)
			expectedMessage.ReplyMarkup = expectedKeyboard

			mockTelegram.EXPECT().GetMe().Return(me, nil)
			mockTelegram.EXPECT().
				GetUserProfilePhotos(tgbotapi.UserProfilePhotosConfig{UserID: me.ID, Limit: 1}).
				Return(tgbotapi.UserProfilePhotos{}, fmt.Errorf("GetUserProfilePhotos error"))

			mockKeyboardButton.EXPECT().Text().Return(expectedProductBtn.Text)
			mockKeyboardButton.EXPECT().Text().Return(expectedReportBrn.Text)
			mockKeyboardButton.EXPECT().Text().Return(expectedSettingBtn.Text)
			handler.WithKeyboardButton(mockKeyboardButton, mockKeyboardButton, mockKeyboardButton)

			msg, err := handler.Execute(context.Background(), user, "")
			assert.NoError(t, err)
			assert.Equal(t, expectedMessage, msg)
		})

		t.Run("сообщение с фотографией", func(t *testing.T) {
			handler, mockKeyboardButton, mockTelegram := beforeTest(t)
			fileID := "cb196684-d6ec-4fb1-a5bf-c87e39193514"
			expectedPhoto := tgbotapi.NewPhoto(user.TelegramID(), tgbotapi.FileID(fileID))
			expectedPhoto.Caption = expectedText
			expectedPhoto.ReplyMarkup = expectedKeyboard

			mockTelegram.EXPECT().GetMe().Return(me, nil)
			mockTelegram.EXPECT().
				GetUserProfilePhotos(tgbotapi.UserProfilePhotosConfig{UserID: me.ID, Limit: 1}).
				Return(tgbotapi.UserProfilePhotos{Photos: [][]tgbotapi.PhotoSize{{{FileID: fileID}}}}, nil)

			mockKeyboardButton.EXPECT().Text().Return(expectedProductBtn.Text)
			mockKeyboardButton.EXPECT().Text().Return(expectedReportBrn.Text)
			mockKeyboardButton.EXPECT().Text().Return(expectedSettingBtn.Text)
			handler.WithKeyboardButton(mockKeyboardButton, mockKeyboardButton, mockKeyboardButton)

			msg, err := handler.Execute(context.Background(), user, "")

			assert.NoError(t, err)
			assert.Equal(t, expectedPhoto, msg)
		})
	})
}

func TestAlias(t *testing.T) {
	t.Parallel()

	t.Run("возврщается название команды", func(t *testing.T) {
		handler := NewHandler(nil, nil)

		assert.Equal(t, "start", handler.Alias())
	})
}

func TestDescription(t *testing.T) {
	t.Parallel()

	t.Run("возврщается описание команды", func(t *testing.T) {
		handler := NewHandler(nil, nil)

		assert.Equal(t, "Показать привестивие", handler.Description())
	})
}
