package internal

import (
	"fmt"
	"log"
	"os"

	telegrambotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/IASamoylov/tg_calories_observer/internal/api"
	"github.com/IASamoylov/tg_calories_observer/internal/pkg/graceful"
	multicloser "github.com/IASamoylov/tg_calories_observer/internal/pkg/multi_closer"
	simpleserver "github.com/IASamoylov/tg_calories_observer/internal/pkg/simple_server"
)

func init() {
	// TODO pgx scan with allow unknown columns
}

// InitControllers initializes REST handlers
func (app *app) InitControllers() *app {
	app.controllers = api.GetHandlers()

	return app
}

// InitServer initializes the web server to provide access via REST
func (app *app) InitServer(port string) *app {
	host := fmt.Sprintf(":%s", port)
	handlers := api.GetHandlers()

	app.httpServer = simpleserver.NewHTTPServer(host, handlers...)
	app.closer.Add(app.httpServer)

	return app
}

// InitPgxConnection initializes the pgx driver to connect to postgres
func (app *app) InitPgxConnection() *app {
	return app
}

// InitExternalClientsIfNotSet initializes external services if they was not overridden for integration tests
func (app *app) InitExternalClientsIfNotSet() *app {
	if app.externalClients.TelegramBotAPI == nil {
		api, err := telegrambotapi.NewBotAPI("token")
		if err != nil {
			log.Panicf("an error occurred when creating a telegram client API: %s", err.Error())
		}

		app.externalClients.TelegramBotAPI = api
	}

	return app
}

// ApplyOverridesExtermalClient overrides to be able to test the application in isolation from other systems
func (app *app) ApplyOverridesExtermalClient(overrides ...OverrideExtermalClient) *app {
	for _, override := range overrides {
		app = override(app)
	}

	return app
}

// InitGracefulShutdown initializes the graceful shutdown for the application
func (app *app) InitGracefulShutdown(signals ...os.Signal) *app {
	graceful.Shutdown(multicloser.GetGlobalCloser(), signals...)

	return app
}

// WithTelegramAPI creates service with specific telegram API client
func WithTelegramAPI(ctor func(token string) *telegrambotapi.BotAPI) OverrideExtermalClient {
	return func(app *app) *app {
		app.externalClients.TelegramBotAPI = ctor("token")

		return app
	}
}
