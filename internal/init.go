package internal

import (
	"fmt"
	"net/http"
	"time"

	telegramlogger "github.com/IASamoylov/tg_calories_observer/internal/pkg/logger/cores/telegram"

	"github.com/IASamoylov/tg_calories_observer/internal/config"

	"github.com/IASamoylov/tg_calories_observer/internal/clients/telegram"

	commandrouter "github.com/IASamoylov/tg_calories_observer/internal/app/services/command_router"

	"github.com/IASamoylov/tg_calories_observer/internal/utils/types"

	"github.com/IASamoylov/tg_calories_observer/internal/app/comands/help"

	"github.com/IASamoylov/tg_calories_observer/internal/app/comands/start"

	productkeyboard "github.com/IASamoylov/tg_calories_observer/internal/app/keybords/product"

	"github.com/IASamoylov/tg_calories_observer/internal/app/inline_keyboards/product"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/logger"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/crypto"

	"github.com/IASamoylov/tg_calories_observer/internal/infrastructure/database"

	multicloser "github.com/IASamoylov/tg_calories_observer/internal/pkg/multi_closer"

	"github.com/jackc/pgx/v5/pgxpool"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	debugv1 "github.com/IASamoylov/tg_calories_observer/internal/api/debug/v1"
	telegramv1 "github.com/IASamoylov/tg_calories_observer/internal/api/telegram/v1"
	configdebug "github.com/IASamoylov/tg_calories_observer/internal/config/debug"
	simpleserver "github.com/IASamoylov/tg_calories_observer/internal/pkg/simple_server"
)

func init() {
	// TODO pgx scan with allow unknown columns
}

// InitCryptor инцилизирует шифровальщик, который позволяет скрыть персональные данные.
func (app *App) InitCryptor() *App {
	app.cryptor = crypto.NewCryptor(app.Cfg.CryptorKeys)

	return app
}

// InitControllers инцилизирует ручки сервиса.
func (app *App) InitControllers() *App {
	app.controllers.debug = debugv1.NewController()
	app.controllers.telegram = telegramv1.NewController(
		app.services.commandRouter,
		//app.services.keyboardRouter,
	)

	return app
}

// InitServer инцилизирует REST пути и связывает их с ручками сервиса.
// Зависит от вызова метода InitControllers.
func (app *App) InitServer() *App {
	host := fmt.Sprintf(":%s", app.Cfg.Port)

	// optimizing the use of yandex cloud resources
	var apiPrefix string
	if configdebug.Version == configdebug.BetaVersion {
		apiPrefix = fmt.Sprintf("%s/", configdebug.BetaVersion)
	}

	app.httpServer = simpleserver.
		NewHTTPServer(host, apiPrefix).
		Register(http.MethodGet, "api/v1/debug", app.controllers.debug.V1GetServiceInfo).
		Register(http.MethodPost, "api/v1/telegram/updates", app.controllers.telegram.V1WebhookUpdates)
	app.closer.Add(app.httpServer)

	return app
}

// InitPgxConnection инцилизирует pgx для подключения к базе postgres.
func (app *App) InitPgxConnection() *App {
	pool, err := pgxpool.New(app.ctx, app.Cfg.Postgres.Conn())
	if err != nil {
		logger.Panicf("an error occurred when creating a pool of connections to the database: %s", err.Error())
	}
	app.closer.Add(multicloser.NewIOCloserWrap(func() error {
		pool.Close()

		return nil
	}))

	if err = pool.Ping(app.ctx); err != nil {
		logger.Panicf("an error occurred when picking the database: %s", err.Error())
	}

	app.pool = pool

	return app
}

// InitRepositories инцилизирует репозитории для работы с таблицами базы данных.
// Зависит от вызова методов InitPgxConnection и InitCryptor.
func (app *App) InitRepositories() *App {
	app.repositories = repositories{
		user: database.NewSecurityUserRepository(database.NewUserRepository(app.pool), app.cryptor),
	}

	return app
}

// InitExternalClientsConnIfNotSet инцилизирует подключения к внешним сервис (API)
// под средством протоколов http/gRPC и других. Могут быть переопределены на старте приложения.
func (app *App) InitExternalClientsConnIfNotSet() *App {
	if app.externalClients.telegramBotAPI == nil {
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.MaxIdleConns = 100
		transport.MaxConnsPerHost = 100
		transport.MaxIdleConnsPerHost = 100
		httpClient := &http.Client{
			Timeout:   10 * time.Second,
			Transport: transport,
		}

		api, err := tgbotapi.NewBotAPIWithClient(app.Cfg.Telegram.Token, tgbotapi.APIEndpoint, httpClient)
		if err != nil {
			logger.Panicf("an error occurred when creating a telegram client API: %s", err.Error())
		}

		api.Debug = true
		app.externalClients.telegramBotAPI = api

		// только для полноценной интеграции с телеграммом будет подключен дополнительный логгер
		logger.SetLogger(logger.NewDefault(telegramlogger.NewChannelErrorLoggerCore(
			app.Cfg.Telegram.Support,
			app.externalClients.telegramBotAPI)))
	}

	return app
}

// InitClients инцлизиирует клиенты для взаимодействия с внешними API. Клиенты предоставляют удобный
// функционал для взаимодействия с API. Зависит от вызова метода InitExternalClientsConnIfNotSet.
func (app *App) InitClients() *App {
	app.clients.telegram = telegram.NewClient(app.externalClients.telegramBotAPI)

	return app
}

// InitCommands инлизирует команды доступные боту
func (app *App) InitCommands() *App {
	app.commands.help = help.NewHandler()
	app.commands.start = start.NewHandler(app.commands.help, app.externalClients.telegramBotAPI)

	return app
}

// InitInlineKeyboardButtons инлизирует inline кнопки для упрощения ввода пользователя
func (app *App) InitInlineKeyboardButtons() *App {
	app.inlineKeyboardButtons.addProduct = product.NewAddProductInlineButton(nil)
	app.inlineKeyboardButtons.editProduct = product.NewEditProductInlineButton(nil)
	app.inlineKeyboardButtons.removeProduct = product.NewRemoveProductInlineButton(nil)
	app.inlineKeyboardButtons.getProduct = product.NewGetProductInlineButton(nil)

	return app
}

// InitKeyboards инлизирует кастомную клавиатуру для упрощения ввода пользователя
func (app *App) InitKeyboards() *App {
	app.keyboard.product = productkeyboard.NewButton(
		app.inlineKeyboardButtons.addProduct,
		app.inlineKeyboardButtons.editProduct,
		app.inlineKeyboardButtons.removeProduct,
		app.inlineKeyboardButtons.getProduct,
	)

	app.commands.start.WithKeyboardButton(app.keyboard.product)

	return app
}

// InitServices инцилизирует сервисы, адаптеры.
func (app *App) InitServices() *App {
	app.services.commandRouter = commandrouter.NewCommandRouter(
		app.clients.telegram,
		app.repositories.user,
		app.commands.help,
		app.commands.start,
		app.commands.help,
	)
	//app.services.keyboardRouter = routers.NewKeyboardRouter(
	//	app.externalClients.telegramBotAPI,
	//	app.repositories.user,
	//	app.keyboard.product,
	//)

	return app
}

// ApplyOverrides переопределяет конфигурацию приложения для возможности изоляционного тестирования (e2e) сервиса.
func (app *App) ApplyOverrides(overrides ...AppOverrides) *App {
	for _, override := range overrides {
		app = override(app)
	}

	return app
}

// WithTelegramAPI переопределяет клиент для взаимодействия с telegram API. Для локально разработки используется
// преобразование long pooling в web hook
func WithTelegramAPI(ctor func(cfg config.App) types.Telegram) AppOverrides {
	return func(app *App) *App {
		app.externalClients.telegramBotAPI = ctor(app.Cfg)

		return app
	}
}

// WithCustomPort переопределеяет порт для запуск приложения
func WithCustomPort(port string) AppOverrides {
	return func(app *App) *App {
		app.Cfg.Port = port

		return app
	}
}
