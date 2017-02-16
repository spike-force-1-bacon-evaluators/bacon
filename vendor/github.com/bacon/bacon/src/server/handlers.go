package server

import (
	"fmt"
	"net/http"

	"github.com/bacon/bacon/src/logger"
	"github.com/gorilla/mux"
)

// Handlers its all here...
func handlers() *mux.Router {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(handlerNotFound)
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	r.HandleFunc("/", handlerRoot)
	r.HandleFunc("/ping", handlerPing)
	r.HandleFunc("/map", handlerMap)
	return r
}

// Page not found handler
func handlerNotFound(w http.ResponseWriter, r *http.Request) {
	logger.LogRequest(r.RemoteAddr, r.Method, r.URL.String())
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 - Page Not Found"))
}

// Handler '/'
func handlerRoot(w http.ResponseWriter, r *http.Request) {
	logger.LogRequest(r.RemoteAddr, r.Method, r.URL.String())
	indextmpl(w)
}

// Handler 'ping'. For testing only
func handlerPing(w http.ResponseWriter, r *http.Request) {
	logger.LogRequest(r.RemoteAddr, r.Method, r.URL.String())
	fmt.Fprint(w, "PONG")
}

func handlerMap(w http.ResponseWriter, r *http.Request) {
	logger.LogRequest(r.RemoteAddr, r.Method, r.URL.String())
	maptmpl(w)
}
