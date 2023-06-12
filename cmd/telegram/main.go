package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/graceful"
	multicloser "github.com/IASamoylov/tg_calories_observer/internal/pkg/multi_closer"
	serverready "github.com/IASamoylov/tg_calories_observer/internal/pkg/server_ready"
)

func main() {
	log.Println("start")

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "9090"
	}

	multicloser.AddGlobal(serverready.NewHTTPServer(fmt.Sprintf(":%s", port)).Run())
	graceful.Shutdown(multicloser.GetGlobalCloser(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	multicloser.WaitGlobal()
	log.Print("Server Stopped")
}
