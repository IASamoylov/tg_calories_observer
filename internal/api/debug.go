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
	AppName        string `json:"app_name"`
	GithubSHA      string `json:"github_sha"`
	GithubSHAShort string `json:"github_sha_short"`
	Version        string `json:"version"`
	BuildedAt      string `json:"builded_at"`
	Message        string `json:"message"`
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
		Message:        "dev-002",
	})
	if err != nil {
		log.Println(fmt.Sprintf("an error occurred while writing response: %s :)", err))
	}
	w.WriteHeader(http.StatusOK)
}
