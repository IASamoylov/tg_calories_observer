package internal

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/types"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	debugv1 "github.com/IASamoylov/tg_calories_observer/internal/api/debug/v1"
	telegramv1 "github.com/IASamoylov/tg_calories_observer/internal/api/telegram/v1"
	"github.com/IASamoylov/tg_calories_observer/internal/clients/telegram"
	config "github.com/IASamoylov/tg_calories_observer/internal/config/debug"
	"github.com/IASamoylov/tg_calories_observer/internal/pkg/graceful"
	simpleserver "github.com/IASamoylov/tg_calories_observer/internal/pkg/simple_server"
)

func init() {
	// TODO pgx scan with allow unknown columns
}

// InitControllers initializes REST handlers
func (app *App) InitControllers() *App {
	app.controllers.debug = debugv1.NewController()
	app.controllers.telegram = telegramv1.NewController(app.ExternalClients.telegramBotAPI)

	return app
}

// InitServer initializes the web server to provide access via REST
func (app *App) InitServer(port string) *App {
	host := fmt.Sprintf(":%s", port)

	// optimizing the use of yandex cloud resources
	var apiPrefix string
	if config.Version == config.BetaVersion {
		apiPrefix = fmt.Sprintf("%s/", config.BetaVersion)
	}

	app.httpServer = simpleserver.
		NewHTTPServer(host, apiPrefix).
		Register(http.MethodGet, "api/v1/debug", app.controllers.debug.V1GetServiceInfo).
		Register(http.MethodPost, "api/v1/telegram/updates", app.controllers.telegram.V1WebhookUpdates)
	app.closer.Add(app.httpServer)

	return app
}

// InitPgxConnection initializes the pgx driver to connect to postgres
func (app *App) InitPgxConnection() *App {
	return app
}

// InitExternalClientsConnIfNotSet initializes external services if they was not overridden for integration tests
func (app *App) InitExternalClientsConnIfNotSet() *App {

	if app.ExternalClients.telegramBotAPI == nil {
		// https://core.telegram.org/bots/webhooks#testing-your-bot-with-updates
		api, err := tgbotapi.NewBotAPI(app.Cfg.Telegram.Token)

		if err != nil {
			log.Panicf("an error occurred when creating a telegram client API: %s", err.Error())
		}

		app.ExternalClients.telegramBotAPI = api
	}

	return app
}

// InitClients initializes external services if they was not overridden for integration tests
func (app *App) InitClients() *App {
	app.clients.telegramClient = telegram.NewTelegramClient(app.ExternalClients.telegramBotAPI)

	return app
}

// ApplyOverridesExternalClientConn overrides to be able to test the application in isolation from other systems
func (app *App) ApplyOverridesExternalClientConn(overrides ...OverrideExternalClient) *App {
	for _, override := range overrides {
		app = override(app)
	}

	return app
}

// InitGracefulShutdown initializes the graceful shutdown for the application
func (app *App) InitGracefulShutdown(signals ...os.Signal) *App {
	graceful.Shutdown(app.closer, signals...)

	return app
}

// WithTelegramAPI creates service with specific telegram API client
func WithTelegramAPI(ctor func(token string) types.TelegramBotAPI) OverrideExternalClient {
	return func(app *App) *App {
		app.ExternalClients.telegramBotAPI = ctor(app.Cfg.Telegram.Token)

		return app
	}
}
