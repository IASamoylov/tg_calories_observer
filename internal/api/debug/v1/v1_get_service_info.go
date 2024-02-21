package debug

import (
	"encoding/json"
	"net/http"

	config "github.com/IASamoylov/tg_calories_observer/internal/config/debug"
)

// V1GetServiceInfo returns debug info about service
func (ctr Controller) V1GetServiceInfo(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(map[string]string{
		"app_name":         config.AppName,
		"version":          config.Version,
		"github_sha":       config.GithubSHA,
		"github_sha_short": config.GithubSHAShort,
	})
}
