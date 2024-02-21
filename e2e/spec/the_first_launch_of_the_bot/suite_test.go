//go:build e2e
// +build e2e

package the_first_launch_of_the_bot

import (
	"context"
	"fmt"

	"go.uber.org/mock/gomock"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/IASamoylov/tg_calories_observer/e2e/asserts"
	"github.com/IASamoylov/tg_calories_observer/internal/domain/entity/dto"

	"github.com/IASamoylov/tg_calories_observer/e2e/fixtures"

	"github.com/IASamoylov/tg_calories_observer/e2e/app"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

var testableApp *app.TestableApp

func TestTheFirstLaunchOfTheBot(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Первое знакомство с ботом")
}

var _ = BeforeSuite(func(ctx context.Context) {
	testableApp = fixtures.RunTestableApp(ctx)
})

var _ = Describe("Новый пользователь", func() {
	var user dto.User

	BeforeEach(func(ctx context.Context) {
		user = testableApp.NewUser(GinkgoParallelProcess())
	})

	Describe("отправляет сообщение с командой /start", func() {
		It("пользователь успешно создается", func(ctx context.Context) {
			defer GinkgoRecover()

			testableApp.TelegramExpect().GetMe().Return(tgbotapi.User{ID: GinkgoRandomSeed()}, nil)
			testableApp.TelegramExpect().
				GetUserProfilePhotos(tgbotapi.UserProfilePhotosConfig{UserID: GinkgoRandomSeed(), Limit: 1}).
				Return(tgbotapi.UserProfilePhotos{}, nil)
			testableApp.TelegramExpect().Send(gomock.Any()).Return(tgbotapi.Message{}, nil)

			fixtures.SendMessage(ctx, testableApp, user, "/start")

			asserts.UserCreatedInStorage(ctx, testableApp, user)
		})

		Context("и бот ответчает сообщением которое содержит", func() {
			When("у бота установлена фотография", func() {
				It("приветсвие с картинкой", func(ctx context.Context) {
					defer GinkgoRecover()

					fileID := user.UserName() + "_file"
					botID := GinkgoRandomSeed()
					testableApp.Mocks().Telegram.EXPECT().GetMe().Return(tgbotapi.User{ID: botID}, nil)
					testableApp.Mocks().Telegram.EXPECT().
						GetUserProfilePhotos(tgbotapi.UserProfilePhotosConfig{UserID: GinkgoRandomSeed(), Limit: 1}).
						Return(tgbotapi.UserProfilePhotos{
							Photos: [][]tgbotapi.PhotoSize{
								{{FileID: fmt.Sprint(fileID)}},
							},
						}, nil)

					testableApp.Mocks().Telegram.EXPECT().Send(gomock.Cond(func(x any) bool {
						val, ok := x.(tgbotapi.PhotoConfig)

						return ok &&
							len(val.Caption) != 0 &&
							val.File == tgbotapi.FileID(fileID) &&
							val.BaseChat.ReplyMarkup != nil
					})).Times(1)

					fixtures.SendMessage(ctx, testableApp, user, "/start")
				})
			})
		})
	})

	Describe("отправляет сообщение с командой /help", func() {
		It("бот ответчает сообщением которое содержит информацию по возможностям сервиса", func(ctx context.Context) {
			defer GinkgoRecover()

			expectedMessage := "Вот что я умею:\n\n" +
				"Если у вас есть вопросы или нужна помощь, не стесняйтесь обращаться!"
			testableApp.TelegramExpect().
				Send(tgbotapi.NewMessage(user.TelegramID(), expectedMessage)).
				Return(tgbotapi.Message{}, nil)

			fixtures.SendMessage(ctx, testableApp, user, "/help")
		})
	})
})

var _ = AfterSuite(func(ctx context.Context) {
	fixtures.StopTestableApp(ctx, testableApp)
})
