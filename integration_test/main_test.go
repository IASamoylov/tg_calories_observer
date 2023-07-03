//go:build integration_test
// +build integration_test

package test_integration

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/IASamoylov/tg_calories_observer/integration_test/debug_handler"

	"github.com/IASamoylov/tg_calories_observer/integration_test/global"

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
	t.Run("DebugHandlerSuite", func(t *testing.T) {
		suite.Run(t, &debug_handler.DebugHandlerSuite{GlobalContext: globalContext})
	})
}
