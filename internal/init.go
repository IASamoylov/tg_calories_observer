package internal

import (
	"fmt"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	debugv1 "github.com/IASamoylov/tg_calories_observer/internal/api/debug/v1"
	telegramv1 "github.com/IASamoylov/tg_calories_observer/internal/api/telegram/v1"
	"github.com/IASamoylov/tg_calories_observer/internal/clients/telegram"
	config "github.com/IASamoylov/tg_calories_observer/internal/config/debug"
	"github.com/IASamoylov/tg_calories_observer/internal/pkg/graceful"
	multicloser "github.com/IASamoylov/tg_calories_observer/internal/pkg/multi_closer"
	simpleserver "github.com/IASamoylov/tg_calories_observer/internal/pkg/simple_server"
)

func init() {
	// TODO pgx scan with allow unknown columns
}

// InitControllers initializes REST handlers
func (app *app) InitControllers() *app {
	app.controllers.debug = debugv1.NewController()
	app.controllers.telegram = telegramv1.NewController(app.externalClients.tgbotapi)

	return app
}

// InitServer initializes the web server to provide access via REST
func (app *app) InitServer(port string) *app {
	host := fmt.Sprintf(":%s", port)

	// optimizing the use of yandex cloud resources
	var apiPrefix string
	if config.Version == config.BetaVersion {
		apiPrefix = config.BetaVersion
	}

	app.httpServer = simpleserver.
		NewHTTPServer(host, apiPrefix).
		Register(http.MethodGet, "/api/v1/debug", app.controllers.debug.V1GetServiceInfo).
		Register(http.MethodPost, "/api/v1/telegram/updates", app.controllers.telegram.V1WebhookUpdates)
	app.closer.Add(app.httpServer)

	return app
}

// InitPgxConnection initializes the pgx driver to connect to postgres
func (app *app) InitPgxConnection() *app {
	return app
}

// InitExternalClientsConnIfNotSet initializes external services if they was not overridden for integration tests
func (app *app) InitExternalClientsConnIfNotSet() *app {

	if app.externalClients.tgbotapi == nil {
		// https://core.telegram.org/bots/webhooks#testing-your-bot-with-updates
		api, err := tgbotapi.NewBotAPI(app.cfg.Telegram.Token)

		if err != nil {
			log.Panicf("an error occurred when creating a telegram client API: %s", err.Error())
		}

		app.externalClients.tgbotapi = api
	}

	return app
}

// InitClients initializes external services if they was not overridden for integration tests
func (app *app) InitClients() *app {
	app.clients.telegramClient = telegram.NewTelegramClient(app.externalClients.tgbotapi)

	return app
}

// ApplyOverridesExternalClientConn overrides to be able to test the application in isolation from other systems
func (app *app) ApplyOverridesExternalClientConn(overrides ...OverrideExternalClient) *app {
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
func WithTelegramAPI(ctor func(token string) *tgbotapi.BotAPI) OverrideExternalClient {
	return func(app *app) *app {
		app.externalClients.tgbotapi = ctor(app.cfg.Telegram.Token)

		return app
	}
}
