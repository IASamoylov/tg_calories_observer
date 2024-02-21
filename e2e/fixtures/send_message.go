//go:build e2e
// +build e2e

package fixtures

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/IASamoylov/tg_calories_observer/internal/domain/entity/dto"

	"github.com/IASamoylov/tg_calories_observer/e2e/app"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// SendMessage проверяет доступность сервиса
var SendMessage = func(ctx context.Context, app *app.TestableApp, user dto.User, text string) {
	By(fmt.Sprintf("Пользователь @%s отправляет сообщение %s", user.UserName(), text))

	update := tgbotapi.Update{
		UpdateID: GinkgoParallelProcess(),
		Message: &tgbotapi.Message{
			From: &tgbotapi.User{
				ID:           user.TelegramID(),
				UserName:     user.UserName(),
				FirstName:    user.FirstName(),
				LastName:     user.LastName(),
				LanguageCode: user.Language(),
			},
			Entities: []tgbotapi.MessageEntity{
				{Type: "bot_command", Offset: 0, Length: len(text)},
			},
			Text: text,
		},
	}

	sendMessage := func() (*http.Response, error) {
		data, err := json.Marshal(update)
		if err != nil {
			return nil, err
		}

		host := fmt.Sprintf("%s/v1/telegram/updates", app.Host())
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, host, bytes.NewBuffer(data))
		if err != nil {
			return nil, err
		}

		resp, err := http.DefaultClient.Do(req)

		return resp, err
	}

	Eventually(sendMessage).WithContext(ctx).
		WithTimeout(100 * time.Millisecond).
		Should(HaveHTTPStatus(http.StatusOK))
}
