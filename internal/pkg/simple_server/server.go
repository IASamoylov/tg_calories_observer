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

// NewHTTPServer creates simple server
func NewHTTPServer(host string) *SimpleHTTPServer {
	return &SimpleHTTPServer{
		base: http.Server{
			Addr:              host,
			WriteTimeout:      1 * time.Second,
			ReadHeaderTimeout: 1 * time.Second,
			Handler:           http.NewServeMux(),
		},
	}
}

// Register registers a handler for the path
func (server *SimpleHTTPServer) Register(
	method string,
	path string,
	handler func(writer http.ResponseWriter, req *http.Request)) *SimpleHTTPServer {
	if mux, ok := server.base.Handler.(*http.ServeMux); ok {
		mux.HandleFunc(path, func(writer http.ResponseWriter, req *http.Request) {
			if req.Method != method {
				writer.WriteHeader(http.StatusNotFound)

				return
			}

			handler(writer, req)
		})
	}

	return server
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
