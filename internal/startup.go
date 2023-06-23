package internal

import (
	"log"
	"os"
	"syscall"

	"github.com/IASamoylov/tg_calories_observer/internal/clients/telegram"
	multicloser "github.com/IASamoylov/tg_calories_observer/internal/pkg/multi_closer"
	simpleserver "github.com/IASamoylov/tg_calories_observer/internal/pkg/simple_server"
)

type externalClients struct {
	TelegramBotAPI telegram.BotAPI
}

type app struct {
	// config
	externalClients *externalClients
	// db poll

	// repositories
	// clients
	// services
	// CQRS

	controllers []simpleserver.RegisterHandler
	closer      *multicloser.MultiCloser
	httpServer  *simpleserver.SimpleHTTPServer
}

// OverrideExtermalClient functions to replace an external clients with mocks for integration tests
type OverrideExtermalClient func(app *app) *app

// NewApp creates a new app with all dependecies
func NewApp(port string, overrides ...OverrideExtermalClient) *app {
	app := &app{
		closer:          multicloser.New(),
		externalClients: &externalClients{},
	}

	app.ApplyOverridesExtermalClient(overrides...).
		InitExternalClientsIfNotSet().
		InitPgxConnection().
		InitControllers().
		InitServer(port).
		InitGracefulShutdown(os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	return app
}

func (app *app) Run() {
	app.httpServer.Run()

	app.closer.Wait()

	log.Print("Server Stopped")
}
