package server

import (
	"fmt"
	"log"
	"net/http"
)

const port = "8088"

// Run trigger server initialization
func Run() {
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Now listening on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, handlers()))
}
