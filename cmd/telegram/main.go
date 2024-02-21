package main

import (
	"context"

	app "github.com/IASamoylov/tg_calories_observer/internal"
)

func main() {
	app.NewApp(context.Background()).Run()
}
