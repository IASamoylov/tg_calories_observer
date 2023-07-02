package internal

import (
	"log"
	"os"
	"syscall"

	"github.com/IASamoylov/tg_calories_observer/internal/config"

	debugv1 "github.com/IASamoylov/tg_calories_observer/internal/api/debug/v1"
	telegramv1 "github.com/IASamoylov/tg_calories_observer/internal/api/telegram/v1"
	"github.com/IASamoylov/tg_calories_observer/internal/clients/telegram"
	multicloser "github.com/IASamoylov/tg_calories_observer/internal/pkg/multi_closer"
	simpleserver "github.com/IASamoylov/tg_calories_observer/internal/pkg/simple_server"
	"github.com/IASamoylov/tg_calories_observer/internal/pkg/types"
)

type externalClients struct {
	telegramBotAPI types.TelegramBotAPI
}

type clients struct {
	telegramClient *telegram.Client
}

type controllers struct {
	debug    *debugv1.Controller
	telegram *telegramv1.Controller
}

// App service
type App struct {
	Cfg             *config.App
	ExternalClients *externalClients
	// db poll
	clients clients
	// repositories
	// clients
	// services
	// CQRS

	controllers *controllers
	httpServer  *simpleserver.SimpleHTTPServer
	closer      *multicloser.MultiCloser
}

// OverrideExternalClient functions to replace an external clients with mocks for integration tests
type OverrideExternalClient func(app *App) *App

// NewApp creates a new App with all dependencies
func NewApp(port string, overrides ...OverrideExternalClient) *App {
	app := &App{
		Cfg:             config.NewConfig(),
		closer:          multicloser.New(),
		ExternalClients: &externalClients{},
		controllers:     &controllers{},
	}

	// release resources that have been added globally
	app.closer.Add(multicloser.NewIOCloserWrap(func() error {
		return multicloser.CloseGlobal()
	}))

	app.ApplyOverridesExternalClientConn(overrides...).
		InitExternalClientsConnIfNotSet().
		InitClients().
		InitPgxConnection().
		InitControllers().
		InitServer(port).
		InitGracefulShutdown(os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	return app
}

func (app *App) Run() {
	app.httpServer.Run()

	app.closer.Wait()

	log.Print("Server Stopped")
}
