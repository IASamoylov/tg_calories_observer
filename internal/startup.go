package internal

import (
	"fmt"
	"log"
	"os"
	"syscall"

	api "github.com/IASamoylov/tg_calories_observer/internal/api"
	"github.com/IASamoylov/tg_calories_observer/internal/pkg/graceful"
	multicloser "github.com/IASamoylov/tg_calories_observer/internal/pkg/multi_closer"
	simpleserver "github.com/IASamoylov/tg_calories_observer/internal/pkg/simple_server"
)

type app struct {
	handlers []simpleserver.RegisterHandler

	closer     *multicloser.MultiCloser
	httpServer *simpleserver.SimpleHTTPServer
}

// NewApp creates a new app with all dependecies
func NewApp(port string) *app {
	app := &app{
		closer: multicloser.New(),
	}

	app.initHandlers().
		initServer(port).
		initGracefulShutdown(os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	return app
}

func (app *app) Run() {
	app.httpServer.Run()

	app.closer.Wait()

	log.Print("Server Stopped")
}

func (app *app) initServer(port string) *app {
	host := fmt.Sprintf(":%s", port)
	handlers := api.GetHandlers()

	app.httpServer = simpleserver.NewHTTPServer(host, handlers...)
	app.closer.Add(app.httpServer)

	return app
}

func (app *app) initHandlers() *app {
	app.handlers = api.GetHandlers()

	return app
}

func (app *app) initGracefulShutdown(signals ...os.Signal) *app {
	graceful.Shutdown(multicloser.GetGlobalCloser(), signals...)

	return app
}
