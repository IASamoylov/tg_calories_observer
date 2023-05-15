package serverready

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

const handlerName string = "/ready"

type server struct {
	base http.Server
}

// NewHTTPServer creates simple server with handler /ready, for check after run
func NewHTTPServer(host string) *server {
	mux := http.NewServeMux()
	mux.HandleFunc(handlerName, ready)

	return &server{
		base: http.Server{
			Addr:              host,
			WriteTimeout:      1 * time.Second,
			ReadHeaderTimeout: 1 * time.Second,
			Handler:           mux,
		},
	}
}

// Start starts a new server in goroutine
func (server *server) Run() *server {
	go func() {
		err := server.base.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println(fmt.Sprintf("an error occurred while executing http.Server.ListenAndServe: %s ", err))
		}
	}()

	log.Println(fmt.Sprintf("Server stated on host %s", server.base.Addr))

	return server
}

// Close stops server
func (server *server) Close() error {
	if err := server.base.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("an error occurred while executing http.Server.Shutdown: %s ", err)
	}

	return nil
}

func ready(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}
