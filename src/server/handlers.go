package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handlers its all here...
func handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ping", handlerPing)
	return r
}

// For testing only
func handlerPing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PONG")
}
