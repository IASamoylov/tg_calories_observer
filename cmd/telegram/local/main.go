package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IASamoylov/tg_calories_observer/internal/config"

	"github.com/IASamoylov/tg_calories_observer/internal/utils/types"

	app "github.com/IASamoylov/tg_calories_observer/internal"
	"github.com/IASamoylov/tg_calories_observer/internal/pkg/logger"
	multicloser "github.com/IASamoylov/tg_calories_observer/internal/pkg/multi_closer"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	app.NewApp(context.Background(), app.WithTelegramAPI(func(cfg config.App) types.Telegram {
		api, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)

		if err != nil {
			logger.Panicf("an error occurred when creating a telegram client API: %s", err.Error())
		}

		// converts long polling to webhook integration for local development
		go func() {
			u := tgbotapi.NewUpdate(0)
			u.Timeout = 60

			updates := api.GetUpdatesChan(u)

			ctx, cancel := context.WithCancel(context.Background())

			multicloser.AddGlobal(multicloser.NewIOCloserWrap(func() error {
				cancel()

				return nil
			}))

			for update := range updates {
				select {
				case <-ctx.Done():
					return
				default:
					msg, _ := json.Marshal(update)
					host := fmt.Sprintf("http://localhost:%s/api/v1/telegram/updates", cfg.Port)
					// nolint
					_, err = http.Post(host, "application/json", bytes.NewBuffer(msg))
					if err != nil {
						logger.Errorf("an error occurred when send POST request: %s", err.Error())
					}
				}
			}
		}()

		return api
	})).Run()
}
