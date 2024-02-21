package commandrouter

import (
	"context"
	"errors"
	"testing"

	"github.com/IASamoylov/tg_calories_observer/internal/domain/entity/dto"
	"github.com/IASamoylov/tg_calories_observer/internal/utils/types"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/mock/gomock"
)

var user = dto.NewUser(100500, "test", "", "", "")

func beforeTest(t *testing.T) (*CommandRouter, *Mocktelegram, *MockuserStorage, *MockcommandRouterHandler) {
	mockTelegram := NewMocktelegram(gomock.NewController(t))
	mockUserStorage := NewMockuserStorage(gomock.NewController(t))

	mockHelpCommand := NewMockcommandRouterHandler(gomock.NewController(t))
	mockHelpCommand.EXPECT().Alias().Return("/help").AnyTimes()
	mockHelpCommand.EXPECT().Description().Return("помощь").AnyTimes()

	mockStartCommand := NewMockcommandRouterHandler(gomock.NewController(t))
	mockStartCommand.EXPECT().Alias().Return("/start").AnyTimes()
	mockStartCommand.EXPECT().Description().Return("запуск").AnyTimes()

	router := NewCommandRouter(mockTelegram, mockUserStorage, mockHelpCommand, mockStartCommand, mockHelpCommand)

	return router, mockTelegram, mockUserStorage, mockStartCommand
}

func TestExecute(t *testing.T) {
	t.Parallel()

	t.Run("корректно обработаная паника приведет к сообщению с ошибкой", func(t *testing.T) {
		t.Parallel()

		mockTelegram := NewMocktelegram(gomock.NewController(t))
		router := NewCommandRouter(mockTelegram, nil, nil)

		mockTelegram.EXPECT().SendErr(user.TelegramID(), types.ErrCommon)

		router.Execute(context.Background(), user, "/start", "")
	})

	t.Run("обработка неизвестной команды приведет к сообщению с ошибкой", func(t *testing.T) {
		t.Parallel()

		router, mockTelegram, mockUserStorage, mockStartCmd := beforeTest(t)

		mockUserStorage.EXPECT().Upsert(gomock.Any(), user).Return(nil)

		expectedErr := errors.New("@test к сожалению, ваш запрос '/hello' не распознан. " +
			"Прошу прощения за неудобства. Продолжайте следить за инструкциями /help и " +
			"оставайтесь на связи для дальнейших указаний.")
		mockTelegram.EXPECT().SendErr(user.TelegramID(), expectedErr)
		mockStartCmd.EXPECT().Execute(gomock.Any(), user, "").Times(0)

		router.Execute(context.Background(), user, "/hello", "")
	})

	t.Run("получение ошибки при акутализации данных пользоватя приведет к сообщению с ошибкой", func(t *testing.T) {
		t.Parallel()

		router, mockTelegram, mockUserStorage, mockStartCmd := beforeTest(t)

		mockUserStorage.EXPECT().Upsert(gomock.Any(), user).Return(errors.New("Upsert err"))

		mockTelegram.EXPECT().SendErr(user.TelegramID(), types.ErrCommon)
		mockStartCmd.EXPECT().Execute(gomock.Any(), user, "").Times(0)

		router.Execute(context.Background(), user, "/start", "")
	})

	t.Run("пустое сообщение не будет обработано", func(t *testing.T) {
		t.Parallel()

		router, _, mockUserStorage, mockStartCmd := beforeTest(t)

		mockUserStorage.EXPECT().Upsert(gomock.Any(), user).Return(nil)
		mockStartCmd.EXPECT().Execute(gomock.Any(), user, "").Times(0)

		router.Execute(context.Background(), user, "", "")
	})

	t.Run("получение ошибки при обработки команды приведет к сообщению с ошибкой", func(t *testing.T) {
		t.Parallel()

		router, mockTelegram, mockUserStorage, mockStartCmd := beforeTest(t)

		mockUserStorage.EXPECT().Upsert(gomock.Any(), user).Return(nil)
		mockStartCmd.EXPECT().Execute(gomock.Any(), user, "").Return(nil, errors.New("Execute error"))
		mockTelegram.EXPECT().SendErr(user.TelegramID(), types.ErrCommon)

		router.Execute(context.Background(), user, "/start", "")
	})

	t.Run("команда будет корректно обработана", func(t *testing.T) {
		t.Parallel()

		router, mockTelegram, mockUserStorage, mockStartCmd := beforeTest(t)

		expectedMsg := tgbotapi.NewMessage(user.TelegramID(), "hello")

		mockUserStorage.EXPECT().Upsert(gomock.Any(), user).Return(nil)
		mockStartCmd.EXPECT().Execute(gomock.Any(), user, "args, args").Return(expectedMsg, nil)
		mockTelegram.EXPECT().Send(expectedMsg)

		router.Execute(context.Background(), user, "/start", "args, args")
	})
}

func TestInitMenu(t *testing.T) {
	t.Parallel()
	t.Run("Инцилизирует подсказки для доступных комманд бота", func(t *testing.T) {
		t.Parallel()

		router, mockTelegram, _, _ := beforeTest(t)

		mockTelegram.EXPECT().InitMenu(
			[]tgbotapi.BotCommand{
				{Command: "/start", Description: "запуск"},
				{Command: "/help", Description: "помощь"},
			},
		)

		router.InitMenu()
	})
}
