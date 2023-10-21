package simpleserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/IASamoylov/tg_calories_observer/internal/pkg/logger"
)

// SimpleHTTPServer HTTP server
type SimpleHTTPServer struct {
	apiPrefix string
	base      http.Server
}

// NewHTTPServer создает простой сервер
func NewHTTPServer(host, apiPrefix string) *SimpleHTTPServer {
	return &SimpleHTTPServer{
		apiPrefix: apiPrefix,
		base: http.Server{
			Addr:              host,
			WriteTimeout:      1 * time.Second,
			ReadHeaderTimeout: 1 * time.Second,
			Handler:           http.NewServeMux(),
		},
	}
}

// Register добавляет HTTP обработчик по заявленому пути
func (server *SimpleHTTPServer) Register(
	method string,
	path string,
	handler func(writer http.ResponseWriter, req *http.Request)) *SimpleHTTPServer {
	if mux, ok := server.base.Handler.(*http.ServeMux); ok {
		path = fmt.Sprintf("/%s%s", server.apiPrefix, path)

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

// Run запускает сервер
func (server *SimpleHTTPServer) Run() {
	go func() {
		err := server.base.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Infof("произошла ошибка при ралоте сервера: %s", err)
		}
	}()

	logger.Infof("HTTP сервер запущен на хосте %s", server.base.Addr)
}

// Close останавливат сервер
func (server *SimpleHTTPServer) Close() error {
	return server.base.Shutdown(context.Background())
}
