//go:build e2e
// +build e2e

package global

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/IASamoylov/tg_calories_observer/internal/domain"

	_ "github.com/lib/pq"

	"github.com/pressly/goose/v3"

	"github.com/IASamoylov/tg_calories_observer/internal"
	"github.com/IASamoylov/tg_calories_observer/internal/pkg/logger"
	"github.com/IASamoylov/tg_calories_observer/internal/pkg/types"
	"github.com/IASamoylov/tg_calories_observer/internal/pkg/types/mocks"
	"github.com/golang/mock/gomock"
)

const PORT = "9093"

// Context describes the global context of integration tests with the service inside
type Context struct {
	context.Context
	Host           string
	app            *internal.App
	telegramUserID int64
	// mocks
	TelegramAPIMock *mocks.MockTelegramBotAPI
}

// NewGlobalContext ctor
func NewGlobalContext() *Context {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	telegramAPIMock := mocks.NewMockTelegramBotAPI(gomock.NewController(&testing.T{}))
	global := &Context{
		Context:         ctx,
		Host:            fmt.Sprintf("http://localhost:%s/api", PORT),
		TelegramAPIMock: telegramAPIMock,
		telegramUserID:  1_000_000,
		app: internal.NewApp(ctx, internal.WithTelegramAPI(func(_ string) types.TelegramBotAPI {
			return telegramAPIMock
		})),
	}

	go global.app.Run()

	return global
}

func (c *Context) ApplyMigrations() {
	db, err := sql.Open("postgres", c.app.Cfg.Postgres.Conn())
	if err != nil {
		logger.Panicf("an error occurred when creating a connection to the database: %s", err.Error())
	}

	if err := goose.Up(db, "../migrations"); err != nil {
		logger.Panicf("an error occurred while rolling migrations: %s", err)
	}
}

func (c *Context) ResetMigrations() {
	db, err := sql.Open("postgres", c.app.Cfg.Postgres.Conn())
	if err != nil {
		logger.Panicf("an error occurred when creating a connection to the database: %s", err.Error())
	}

	if err := goose.Reset(db, "../migrations"); err != nil {
		logger.Panicf("an error occurred while rolling migrations: %s", err)
	}
}

func (c *Context) WaitForRun() {
	logger.Infof("Waits for server starts on %s", c.Host)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-c.Done():
				logger.Fatalf("Server is not started")
			default:
				resp, err := http.Get(fmt.Sprintf("%s/v1/debug", c.Host))
				if err == nil && resp.StatusCode == http.StatusOK {
					return
				}
			}
		}
	}()

	wg.Wait()
}

func (c *Context) NextTelegramUserID() domain.TelegramID {
	return domain.TelegramID(atomic.AddInt64(&c.telegramUserID, 10))
}
