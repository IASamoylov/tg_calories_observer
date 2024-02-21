//go:build e2e
// +build e2e

package fixtures

import (
	"context"
	"fmt"
	"time"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/logger"

	"github.com/IASamoylov/tg_calories_observer/e2e/matchers"

	"github.com/IASamoylov/tg_calories_observer/e2e/app"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// RunTestableApp проверяет доступность сервиса
var RunTestableApp = func(ctx context.Context) *app.TestableApp {
	logger.SetLogger(logger.New(GinkgoWriter))

	By("Запуск сервиса")
	testableApp, err := app.NewTestableApp(GinkgoT(), GinkgoRandomSeed(), GinkgoParallelProcess())
	Expect(err).Should(Succeed())

	testableApp.Run()

	Eventually(testableApp.Ping).
		WithContext(ctx).
		WithTimeout(50 * time.Millisecond).
		Should(HaveHTTPBody(matchers.MatchValueAsJSON(map[string]string{
			"version":          "integration",
			"app_name":         "calories-observer-telegram-bot",
			"github_sha":       fmt.Sprintf("%d", GinkgoRandomSeed()),
			"github_sha_short": fmt.Sprintf("%d", GinkgoParallelProcess()),
		})))

	return testableApp
}

// StopTestableApp проверяет, что сервис остановлен все ресурсы освобождены
var StopTestableApp = func(ctx context.Context, app *app.TestableApp) {
	By("Остановка сервиса")
	if app != nil {
		app.Stop()
		Eventually(func(ctx context.Context) error {
			_, err := app.Ping(ctx)

			return err
		}).
			WithContext(ctx).
			WithTimeout(50 * time.Millisecond).
			Should(MatchError(ContainSubstring("connect: connection refused")))
	}
}
