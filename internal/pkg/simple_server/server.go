package simpleserver

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

// SimpleHTTPServer HTTP server
type SimpleHTTPServer struct {
	base http.Server
}

// RegisterHandler describes HTTP handler to receive requests
type RegisterHandler interface {
	GetName() string
	Handle(http.ResponseWriter, *http.Request)
}

// NewHTTPServer creates simple server
func NewHTTPServer(host string, handlers ...RegisterHandler) *SimpleHTTPServer {
	mux := http.NewServeMux()

	for _, handler := range handlers {
		mux.HandleFunc(handler.GetName(), handler.Handle)
	}

	return &SimpleHTTPServer{
		base: http.Server{
			Addr:              host,
			WriteTimeout:      1 * time.Second,
			ReadHeaderTimeout: 1 * time.Second,
			Handler:           mux,
		},
	}
}

// Run starts a new server in goroutine
func (server *SimpleHTTPServer) Run() {
	go func() {
		err := server.base.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println(fmt.Sprintf("an error occurred while executing http.Server.ListenAndServe: %s ", err))
		}
	}()

	log.Println(fmt.Sprintf("HTTP server stated on host %s", server.base.Addr))
}

// Close stops server
func (server *SimpleHTTPServer) Close() error {
	if err := server.base.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("an error occurred while executing http.Server.Shutdown: %s ", err)
	}

	return nil
}
