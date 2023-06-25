package internal

import (
	"log"
	"os"
	"syscall"

	debugv1 "github.com/IASamoylov/tg_calories_observer/internal/api/debug/v1"
	telegramv1 "github.com/IASamoylov/tg_calories_observer/internal/api/telegram/v1"
	"github.com/IASamoylov/tg_calories_observer/internal/clients/telegram"
	multicloser "github.com/IASamoylov/tg_calories_observer/internal/pkg/multi_closer"
	simpleserver "github.com/IASamoylov/tg_calories_observer/internal/pkg/simple_server"
	"github.com/IASamoylov/tg_calories_observer/internal/pkg/types"
)

type externalClients struct {
	tgbotapi types.TelegramBotAPI
}

type clients struct {
	telegramClient *telegram.Client
}

type controllers struct {
	debug    *debugv1.Controller
	telegram *telegramv1.Controller
}

type app struct {
	// config
	externalClients *externalClients
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

// OverrideExtermalClient functions to replace an external clients with mocks for integration tests
type OverrideExtermalClient func(app *app) *app

// NewApp creates a new app with all dependecies
func NewApp(port string, overrides ...OverrideExtermalClient) *app {
	app := &app{
		closer:          multicloser.New(),
		externalClients: &externalClients{},
		controllers:     &controllers{},
	}

	// release resources that have been added globally
	app.closer.Add(multicloser.NewIOCloserWrap(func() error {
		return multicloser.CloseGlobal()
	}))

	app.ApplyOverridesExtermalClientConn(overrides...).
		InitExternalClientsConnIfNotSet().
		InitClients().
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
