package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handlers its all here...
func handlers() *mux.Router {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(handlerNotFound)
	r.HandleFunc("/ping", handlerPing)
	return r
}

// For testing only
// Page not found handler
func handlerNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 - Page Not Found"))
}
func handlerPing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PONG")
}
