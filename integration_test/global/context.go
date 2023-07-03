//go:build integration_test
// +build integration_test

package global

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"

	_ "github.com/lib/pq"

	"github.com/pressly/goose/v3"

	"github.com/IASamoylov/tg_calories_observer/internal"
	"github.com/IASamoylov/tg_calories_observer/internal/pkg/types"
	"github.com/IASamoylov/tg_calories_observer/internal/pkg/types/mocks"
	"github.com/golang/mock/gomock"
)

const PORT = "8080"

// Context describes the global context of integration tests with the service inside
type Context struct {
	context.Context
	Host string
	app  *internal.App

	// mocks
	telegramAPIMock *mocks.MockTelegramBotAPI
}

// NewGlobalContext ctor
func NewGlobalContext() *Context {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Minute)
	telegramAPIMock := mocks.NewMockTelegramBotAPI(gomock.NewController(&testing.T{}))
	global := &Context{
		Context:         ctx,
		Host:            fmt.Sprintf("http://localhost:%s/api", PORT),
		telegramAPIMock: telegramAPIMock,
		app: internal.NewApp(ctx, PORT, internal.WithTelegramAPI(func(_ string) types.TelegramBotAPI {
			return telegramAPIMock
		})),
	}

	go global.app.Run()

	return global
}

func (c *Context) ApplyMigrations() {
	db, err := sql.Open("postgres", c.app.Cfg.Postgres.Conn())
	if err != nil {
		log.Panicf("an error occurred when creating a connection to the database: %s", err.Error())
	}

	if err := goose.Up(db, "../migrations"); err != nil {
		log.Panicf("an error occurred while rolling migrations: %s", err)
	}
}

func (c *Context) ResetMigrations() {
	db, err := sql.Open("postgres", c.app.Cfg.Postgres.Conn())
	if err != nil {
		log.Panicf("an error occurred when creating a connection to the database: %s", err.Error())
	}

	if err := goose.Reset(db, "../migrations"); err != nil {
		log.Panicf("an error occurred while rolling migrations: %s", err)
	}
}

func (c *Context) WaitForRun() {
	log.Println("Waits for server starts")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-c.Done():
				log.Fatalln("Server is not started")
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
