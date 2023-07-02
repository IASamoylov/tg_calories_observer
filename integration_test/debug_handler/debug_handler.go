//go:build integration_test
// +build integration_test

package debug_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/IASamoylov/tg_calories_observer/integration_test/global"
	"github.com/IASamoylov/tg_calories_observer/internal/config/debug"

	"github.com/stretchr/testify/suite"
)

// DebugHandlerSuite a test suite struct that embeds suite.Suite
type DebugHandlerSuite struct {
	suite.Suite
	ctx context.Context
	gc  *global.Context
}

// NewRunnerDebugHandlerSuite creates a new runner
func NewRunnerDebugHandlerSuite(t *testing.T, gc *global.Context) {
	suite.Run(t, &DebugHandlerSuite{gc: gc})
}

// SetupTest method that will be called before each test
func (s *DebugHandlerSuite) SetupTest() {
	s.ctx, _ = context.WithTimeout(s.gc, 500*time.Millisecond)
}

// TestV1GetServiceInfo individual test functions that will be run by the suite
func (s *DebugHandlerSuite) TestV1GetServiceInfo() {
	s.Run("the handle returns information about the service: version, revision number and build time", func() {
		buildTime := time.Now().UTC().Format(time.RFC1123)

		debug.Version = "integration"
		debug.AppName = "calories-observer-telegram-bot"
		debug.GithubSHA = "f616bd7c833e745977d75b8107eeae0173a288ab"
		debug.GithubSHAShort = "f616bd7c"
		debug.BuildTime = buildTime

		req, err := http.NewRequestWithContext(s.ctx, http.MethodGet, fmt.Sprintf("%s/v1/debug", s.gc.Host), nil)
		s.Require().NoError(err, "an error occurred when forming a request to the handle /api/v1/debug")

		resp, err := http.DefaultClient.Do(req)
		s.Require().NoError(err, "request to handle /api/v1/debug completed with error")
		defer resp.Body.Close()

		var debugInfo map[string]string
		s.Require().NoError(json.NewDecoder(resp.Body).Decode(&debugInfo), "the handle returned an unexpected data type")

		s.EqualValues(map[string]string{
			"version":          "integration",
			"app_name":         "calories-observer-telegram-bot",
			"github_sha":       "f616bd7c833e745977d75b8107eeae0173a288ab",
			"github_sha_short": "f616bd7c",
			"build_time":       buildTime,
		}, debugInfo)
	})
}
