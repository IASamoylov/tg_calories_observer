package internal

import (
	"context"
	"os"
	"syscall"

	commandrouter "github.com/IASamoylov/tg_calories_observer/internal/app/services/command_router"

	"github.com/IASamoylov/tg_calories_observer/internal/clients/telegram"

	"github.com/IASamoylov/tg_calories_observer/internal/utils/types"

	"github.com/IASamoylov/tg_calories_observer/internal/app/comands/help"

	"github.com/IASamoylov/tg_calories_observer/internal/app/comands/start"

	"github.com/IASamoylov/tg_calories_observer/internal/app/inline_keyboards/product"
	productkeyboard "github.com/IASamoylov/tg_calories_observer/internal/app/keybords/product"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/graceful"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/crypto"

	"github.com/IASamoylov/tg_calories_observer/internal/infrastructure/database"

	"github.com/IASamoylov/tg_calories_observer/internal/config"

	debugv1 "github.com/IASamoylov/tg_calories_observer/internal/api/debug/v1"
	telegramv1 "github.com/IASamoylov/tg_calories_observer/internal/api/telegram/v1"
	multicloser "github.com/IASamoylov/tg_calories_observer/internal/pkg/multi_closer"
	simpleserver "github.com/IASamoylov/tg_calories_observer/internal/pkg/simple_server"
)

type externalClients struct {
	telegramBotAPI types.Telegram
}

type clients struct {
	telegram *telegram.Client
}

type controllers struct {
	debug    debugv1.Controller
	telegram telegramv1.Controller
}

type repositories struct {
	user database.SecurityUserRepository
}

type services struct {
	commandRouter *commandrouter.CommandRouter
	//keyboardRouter *routers.KeyboardRouter
}

type inlineKeyboardButtons struct {
	// product
	addProduct    *product.AddProductInlineButton
	editProduct   *product.EditProductInlineButton
	removeProduct *product.RemoveProductInlineButton
	getProduct    *product.GetProductInlineButton
}

type command struct {
	start *start.Handler
	help  *help.Handler
}

type keyboard struct {
	// product
	product *productkeyboard.Handler
}

// App ...
type App struct {
	context.Context
	Cfg                   config.App
	cryptor               crypto.Cryptor
	externalClients       externalClients
	clients               clients
	repositories          repositories
	inlineKeyboardButtons inlineKeyboardButtons
	keyboard              keyboard
	commands              command
	services              services
	controllers           controllers
	pool                  *pgxpool.Pool
	httpServer            *simpleserver.SimpleHTTPServer
	closer                *multicloser.MultiCloser
	ctx                   context.Context
}

// OverrideExternalClient функциония позволяющая переопределить компонент системы,
// переопределяет только внешние клиенты. Удобно при использование е2е тестов чтобы
// не использовать реальные внешние сервисы, а моки.
type OverrideExternalClient func(app *App) *App

// NewApp создает новое приеложение инцилизирая все зависимости последовательно
func NewApp(ctx context.Context, overrides ...OverrideExternalClient) *App {
	app := &App{
		ctx:    ctx,
		Cfg:    config.NewConfig(),
		closer: multicloser.New(),
	}

	// регистриуем освобождение глобальные ресурсов
	app.closer.Add(multicloser.GetGlobalCloser())
	app.closer.Add(multicloser.NewIOCloserWrap(logger.Sync))

	app.InitCryptor().
		ApplyOverridesExternalClientConn(overrides...).
		InitExternalClientsConnIfNotSet().
		InitClients().
		InitPgxConnection().
		InitRepositories().
		InitCommands().
		InitInlineKeyboardButtons().
		InitKeyboards().
		InitServices().
		InitControllers().
		InitServer()

	// регистрируем
	graceful.Shutdown(app.closer, os.Interrupt, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	return app
}

// Run запускает сервер
func (app *App) Run() {
	app.services.commandRouter.InitMenu()
	app.httpServer.Run()
	app.closer.Wait()

	logger.Info("Server Stopped")
}
