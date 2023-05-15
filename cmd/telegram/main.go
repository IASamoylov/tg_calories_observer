package main

import (
	"log"
	"os"
	"syscall"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/gracefull"
	multicloser "github.com/IASamoylov/tg_calories_observer/internal/pkg/multi_closer"
	serverready "github.com/IASamoylov/tg_calories_observer/internal/pkg/server_ready"
)

func main() {
	multicloser.AddGlobal(serverready.NewHTTPServer(":9090").Run())
	gracefull.Shutdown(multicloser.GetGlobalCloser(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	multicloser.WaitGlobal()
	log.Print("Server Stopped")
}
