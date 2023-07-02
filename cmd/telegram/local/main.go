package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/types"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	app "github.com/IASamoylov/tg_calories_observer/internal"
	multicloser "github.com/IASamoylov/tg_calories_observer/internal/pkg/multi_closer"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "9090"
	}

	app.NewApp(port, app.WithTelegramAPI(func(token string) types.TelegramBotAPI {
		api, err := tgbotapi.NewBotAPI(token)

		if err != nil {
			log.Panicf("an error occurred when creating a telegram client API: %s", err.Error())
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
					msg, err := json.Marshal(update)
					log.Println(string(msg))
					if err != nil {
						log.Panicf("an error occurred when marshling a telegram message: %s", err.Error())
					}

					host := fmt.Sprintf("http://localhost:%s/api/v1/telegram/updates", port)
					// nolint
					_, err = http.Post(host, "application/json", bytes.NewBuffer(msg))
					if err != nil {
						log.Panicf("an error occurred when send POST request: %s", err.Error())
					}
				}
			}
		}()

		return api
	})).Run()
}
