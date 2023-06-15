package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	debug_config "github.com/IASamoylov/tg_calories_observer/internal/config/debug"
)

type debugHandler struct {
}

type debugMessage struct {
	AppName        string
	GithubSHA      string
	GithubSHAShort string
	Version        string
	BuildedAt      string
}

func (handler debugHandler) GetName() string {
	return "/debug"
}

func (handler debugHandler) Handle(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(debugMessage{
		AppName:        debug_config.AppName,
		Version:        debug_config.Version,
		GithubSHA:      debug_config.GithubSHA,
		GithubSHAShort: debug_config.GithubSHAShort,
		BuildedAt:      debug_config.BuildedAt,
	})
	if err != nil {
		log.Println(fmt.Sprintf("an error occurred while writing response: %s :)", err))
	}
	w.WriteHeader(http.StatusOK)
}
