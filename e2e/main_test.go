//go:build e2e
// +build e2e

package test_integration

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/IASamoylov/tg_calories_observer/e2e/debug_handler"

	"github.com/IASamoylov/tg_calories_observer/e2e/global"

	"github.com/IASamoylov/tg_calories_observer/internal/config"
	"github.com/IASamoylov/tg_calories_observer/internal/config/debug"
)

var globalContext *global.Context

func TestMain(m *testing.M) {
	debug.Version = "integration"
	debug.AppName = os.Getenv("APP_NAME")
	config.Path = "../config"

	globalContext = global.NewGlobalContext()
	globalContext.ApplyMigrations()
	globalContext.WaitForRun()
	code := m.Run()
	globalContext.ResetMigrations()

	os.Exit(code)
}

func TestIntegration(t *testing.T) {
	t.Parallel()

	t.Run("DebugHandlerSuite", func(t *testing.T) {
		t.Parallel()

		suite.Run(t, &debug_handler.DebugHandlerSuite{GlobalContext: globalContext})
	})
}
