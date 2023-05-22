package main

import (
	"log"
	"os"
	"syscall"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/graceful"
	multicloser "github.com/IASamoylov/tg_calories_observer/internal/pkg/multi_closer"
	serverready "github.com/IASamoylov/tg_calories_observer/internal/pkg/server_ready"
)

func main() {
	log.Println("start")
	multicloser.AddGlobal(serverready.NewHTTPServer(":9090").Run())
	graceful.Shutdown(multicloser.GetGlobalCloser(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	multicloser.WaitGlobal()
	log.Print("Server Stopped")
}
