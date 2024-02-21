package graceful

import (
	"context"
	"io"
	"os"
	"os/signal"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/logger"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks io Closer

// Shutdown закрывает неуправляемый ресурс при получения сигнала об остановке приложения
func Shutdown(ctx context.Context, closer io.Closer, signals ...os.Signal) {
	if len(signals) == 0 {
		logger.Info("не перечсилены сигналы при срабатывание, которых необходимо остановить обработку")

		return
	}

	done := make(chan os.Signal, 1)

	go func() {
		defer close(done)
		signal.Notify(done, signals...)
		select {
		// Остановка на основание сигнала из системы
		case <-done:
		// Остановка на основание отмены контекста
		case <-ctx.Done():
		}

		logger.Info("получени сигнал о закрытие приложения, все неуправляемые ресурсы будут остановлены")
		signal.Stop(done)
		if err := closer.Close(); err != nil {
			logger.Warn(err.Error())
		}
	}()
}
