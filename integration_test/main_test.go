//go:build integration_test
// +build integration_test

package test_integration

import (
	"os"
	"testing"

	"github.com/IASamoylov/tg_calories_observer/integration_test/debug_handler"

	"github.com/IASamoylov/tg_calories_observer/integration_test/global"

	"github.com/IASamoylov/tg_calories_observer/internal/config"
	"github.com/IASamoylov/tg_calories_observer/internal/config/debug"
)

var globalContext *global.Context

func TestMain(m *testing.M) {
	debug.Version = "integration"
	debug.AppName = "calories-observer-telegram-bot"
	config.Path = "../config"

	globalContext = global.NewGlobalContext()
	globalContext.ApplyMigrations()
	globalContext.WaitForRun()
	os.Exit(m.Run())
}

func TestIntegration(t *testing.T) {
	debug_handler.NewRunnerDebugHandlerSuite(t, globalContext)
}
