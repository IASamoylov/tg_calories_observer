package internal

import (
	"context"
	"os"
	"syscall"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/crypto"

	"github.com/IASamoylov/tg_calories_observer/internal/infrastructure/database"

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
	telegramClient telegram.Client
}

type controllers struct {
	debug    debugv1.Controller
	telegram telegramv1.Controller
}

type queries struct {
}

type repositories struct {
	user         database.UserRepository
	securityUser database.SecurityUserRepository
}

type services struct {
}

// App service
type App struct {
	context.Context
	Cfg             config.App
	cryptor         crypto.Cryptor
	externalClients externalClients
	clients         clients
	repositories    repositories
	services        services
	//commands        commands
	queries     queries
	controllers controllers
	pool        *pgxpool.Pool
	httpServer  *simpleserver.SimpleHTTPServer
	closer      *multicloser.MultiCloser
	ctx         context.Context
}

// OverrideExternalClient functions to replace an external clients with mocks for integration tests
type OverrideExternalClient func(app *App) *App

// NewApp creates a new App with all dependencies
func NewApp(ctx context.Context, overrides ...OverrideExternalClient) *App {
	app := &App{
		ctx:    ctx,
		Cfg:    config.NewConfig(),
		closer: multicloser.New(),
	}

	// release resources that have been added globally
	app.closer.Add(multicloser.GetGlobalCloser())
	app.closer.Add(multicloser.NewIOCloserWrap(logger.Sync))

	app.InitCryptor().
		ApplyOverridesExternalClientConn(overrides...).
		InitExternalClientsConnIfNotSet().
		InitClients().
		InitPgxConnection().
		InitRepositories().
		InitCommand().
		InitQueries().
		InitMenu().
		InitServices().
		InitControllers().
		InitServer().
		InitGracefulShutdown(os.Interrupt, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	logger.Info("hello world")
	logger.Error("hello world")
	return app
}

// Run starts server
func (app *App) Run() {
	app.httpServer.Run()
	app.closer.Wait()

	logger.Info("Server Stopped")
}
