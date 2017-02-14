package logger

import (
	"bytes"
	"log"
	"net/url"
	"os"
	"testing"
)

func TestLogRequest(t *testing.T) {

	var buf bytes.Buffer

	// Set log output from to buffer
	log.SetOutput(&buf)

	// Input variables
	remoteAddr := "http://127.0.0.1"
	method := "GET"
	url := &url.URL{Path: "/"}

	LogRequest(remoteAddr, method, url.String())

	// Reset log output redirection
	log.SetOutput(os.Stderr)

	// Remove date and trailing char
	output := buf.String()[20 : len(buf.String())-1]
	expected := "http://127.0.0.1 GET /"

	if output != expected {
		t.Errorf("Expected: '%s'. Received: '%s'", expected, output)
	}
}
