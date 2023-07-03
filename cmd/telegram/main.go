package main

import (
	"context"
	"os"

	app "github.com/IASamoylov/tg_calories_observer/internal"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "9090"
	}

	app.NewApp(context.Background(), port).Run()
}
