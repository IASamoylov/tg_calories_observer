package main

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

func main() {
	log.Println("start")

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "9090"
	}

	multicloser.AddGlobal(simpleserver.NewHTTPServer(
		fmt.Sprintf(":%s", port),
		api.GetHandlers()...,
	).Run())
	graceful.Shutdown(multicloser.GetGlobalCloser(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	multicloser.WaitGlobal()
	log.Print("Server Stopped")
}
