//go:build e2e
// +build e2e

package debug

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/IASamoylov/tg_calories_observer/e2e/global"
	"github.com/IASamoylov/tg_calories_observer/internal/config/debug"

	"github.com/stretchr/testify/suite"
)

// Suite a test suite struct that embeds suite.Suite
type Suite struct {
	suite.Suite
	ctx           context.Context
	GlobalContext *global.Context
}

// SetupTest method that will be called before each test
func (s *Suite) SetupTest() {
	s.ctx, _ = context.WithTimeout(s.GlobalContext, 500*time.Millisecond)
}

// TestV1GetServiceInfo individual test functions that will be run by the suite
func (s *Suite) TestV1GetServiceInfo() {
	s.Run("ручка возвращает информацию о текущией версии сервиса", func() {
		buildTime := time.Now().UTC().Format(time.RFC1123)

		debug.Version = "integration"
		debug.AppName = "calories-observer-telegram-bot"
		debug.GithubSHA = "f616bd7c833e7423454365d75b8107eeae0173a288ab"
		debug.GithubSHAShort = "f616bd7c"
		debug.BuildTime = buildTime

		req, err := http.NewRequestWithContext(s.ctx, http.MethodGet, fmt.Sprintf("%s/v1/debug", s.GlobalContext.Host), nil)
		s.Require().NoError(err)

		resp, err := http.DefaultClient.Do(req)
		s.Require().NoError(err)
		defer resp.Body.Close()

		var debugInfo map[string]string
		s.Require().NoError(json.NewDecoder(resp.Body).Decode(&debugInfo))

		s.EqualValues(map[string]string{
			"version":          "integration",
			"app_name":         "calories-observer-telegram-bot",
			"github_sha":       "f616bd7c833e7423454365d75b8107eeae0173a288ab",
			"github_sha_short": "f616bd7c",
			"build_time":       buildTime,
		}, debugInfo)
	})

	s.Run("вызов не существующей ручки приводит к поулчение 404 NotFound", func() {
		req, err := http.NewRequestWithContext(s.ctx, http.MethodPost, fmt.Sprintf("%s/v1/debug", s.GlobalContext.Host), nil)
		s.Require().NoError(err, "an error occurred when forming a request to the handle /api/v1/debug")

		resp, err := http.DefaultClient.Do(req)
		s.Require().NoError(err, "request to handle /api/v1/debug completed with error")
		s.Assert().Equal(http.StatusNotFound, resp.StatusCode)
	})
}
