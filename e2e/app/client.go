//go:build e2e
// +build e2e

package app

import (
	"context"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	multicloser "github.com/IASamoylov/tg_calories_observer/internal/pkg/multi_closer"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/IASamoylov/tg_calories_observer/internal/domain/entity/dto"

	"github.com/IASamoylov/tg_calories_observer/internal/config"
	"github.com/IASamoylov/tg_calories_observer/internal/config/debug"
	"go.uber.org/mock/gomock"

	"github.com/IASamoylov/tg_calories_observer/internal"
	"github.com/IASamoylov/tg_calories_observer/internal/utils/types"
)

const (
	Version = "integration"
	AppName = "calories-observer-telegram-bot"
)

func NewTestableApp(t gomock.TestReporter, seed int64, process int) (*TestableApp, error) {
	debug.Version = Version
	debug.AppName = AppName
	debug.GithubSHA = fmt.Sprintf("%d", seed)
	debug.GithubSHAShort = fmt.Sprintf("%d", process)
	config.Path = "../../../config"

	mocks := &Mocks{
		Telegram: types.NewMockTelegram(gomock.NewController(t)),
	}

	port := fmt.Sprintf("%d", 8080+process)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

	app := internal.NewApp(ctx,
		internal.WithTelegramAPI(func(cfg config.App) types.Telegram { return mocks.Telegram }),
		internal.WithCustomPort(port),
	)
	pool, err := pgxpool.New(ctx, app.Cfg.Postgres.Conn())
	if err == nil {
		multicloser.AddGlobal(multicloser.NewIOCloserWrap(func() error {
			pool.Close()
			return nil
		}))
	}

	return &TestableApp{
		mocks:      mocks,
		app:        app,
		pool:       pool,
		telegramID: seed * int64(process),
		host:       fmt.Sprintf("http://localhost:%s/api", port),
		cancel:     cancel,
	}, err
}

// Mocks моки, которые используются как заглушки для внешних сервисов
type Mocks struct {
	Telegram *types.MockTelegram
}

// TestableApp описывает приложение, которое необходимо протестировать
type TestableApp struct {
	app        *internal.App
	mocks      *Mocks
	host       string
	telegramID int64
	cancel     context.CancelFunc
	pool       types.PgxPool
}

func (app *TestableApp) Pool() types.PgxPool {
	return app.pool
}

func (app *TestableApp) App() *internal.App {
	return app.app
}

func (app *TestableApp) Host() string {
	return app.host
}

func (app *TestableApp) Run() {
	go func() {
		app.app.Run()
	}()
}

func (app *TestableApp) Mocks() *Mocks {
	return app.mocks
}

func (app *TestableApp) TelegramExpect() *types.MockTelegramMockRecorder {
	return app.mocks.Telegram.EXPECT()
}

func (app *TestableApp) NextTelegramID() int64 {
	return atomic.AddInt64(&app.telegramID, 1)
}

func (app *TestableApp) Ping(ctx context.Context) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/v1/debug", app.host), nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)

	return resp, err
}

func (app *TestableApp) NewUser(process int) dto.User {
	id := app.NextTelegramID()

	return dto.NewUser(
		id,
		fmt.Sprintf("test_user_%d_%d", id, process),
		fmt.Sprintf("first_name_%d_%d", id, process),
		fmt.Sprintf("last_name_%d_%d", id, process),
		"ru",
	)
}

func (app *TestableApp) Stop() {
	app.cancel()
}
