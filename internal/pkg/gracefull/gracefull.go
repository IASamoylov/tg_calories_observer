package gracefull

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
)

//go:generate mockgen -destination=mocks/mocks_signal.go -package=mocks os Signal
//go:generate mockgen -destination=mocks/mocks_closer.go -package=mocks io Closer

// Shutdown closes the resource after receiving a signal about the end of the application
func Shutdown(closer io.Closer, signals ...os.Signal) {
	if len(signals) == 0 {
		log.Println("the signal to stop the application is not set")
	}

	go func() {
		done := make(chan os.Signal, 1)
		defer close(done)
		signal.Notify(done, signals...)
		<-done
		log.Println("an application shutdown signal was received")
		signal.Stop(done)
		if err := closer.Close(); err != nil {
			log.Println(fmt.Sprintf("an error occurred when closing the resource: %s", err))
		}
	}()
}
