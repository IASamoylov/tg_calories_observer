package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type readyHandler struct {
}

func (handler readyHandler) GetName() string {
	return "/ready"
}

func (handler readyHandler) Handle(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(struct{ Message string }{
		Message: "[BOT] Web server started",
	})
	if err != nil {
		log.Println(fmt.Sprintf("an error occurred while writing response: %s :)", err))
	}
	w.WriteHeader(http.StatusOK)
}
