package graceful

import (
	"io"
	"os"
	"os/signal"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/logger"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks io Closer

var done = make(chan os.Signal, 1)

// Shutdown closes the resource after receiving a signal about the end of the application
func Shutdown(closer io.Closer, signals ...os.Signal) {
	if len(signals) == 0 {
		logger.Info("the signal to stop the application is not set")

		return
	}

	go func() {
		defer close(done)
		signal.Notify(done, signals...)
		<-done
		logger.Info("an application shutdown signal was received")
		signal.Stop(done)
		if err := closer.Close(); err != nil {
			logger.Infof("an error occurred when closing the resource: %s", err)
		}
	}()
}
