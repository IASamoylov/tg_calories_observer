package api

import (
	simpleserver "github.com/IASamoylov/tg_calories_observer/internal/pkg/simple_server"
)

// GetHandlers returns a list of all handles
func GetHandlers() []simpleserver.RegisterHandler {
	return []simpleserver.RegisterHandler{readyHandler{}, debugHandler{}}
}
