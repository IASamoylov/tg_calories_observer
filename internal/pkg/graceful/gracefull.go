package graceful

import (
	"io"
	"os"
	"os/signal"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/logger"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks io Closer

var done = make(chan os.Signal, 1)

// Shutdown закрывает неуправляемый ресурс при получения сигнала об остановке приложения
func Shutdown(closer io.Closer, signals ...os.Signal) {
	if len(signals) == 0 {
		logger.Info("не перечсилены сигналы при срабатывание, которых необходимо остановить обработку")

		return
	}

	go func() {
		defer close(done)
		signal.Notify(done, signals...)
		<-done
		logger.Info("получени сигнал о закрытие приложения, все не управляемые ресурсы будут остановлены")
		signal.Stop(done)
		if err := closer.Close(); err != nil {
			logger.Warn(err.Error())
		}
	}()
}
