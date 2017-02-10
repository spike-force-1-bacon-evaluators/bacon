package server

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test "ping" endpoint
func TestHandlerPing(t *testing.T) {

	// Prepare GET request to ping endpoint
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		log.Fatal("Failed to prepare request: ", err)
	}

	// record the response
	rr := httptest.NewRecorder()

	// Start test server prepared to receive request on 'ping' endpoint
	handler := http.HandlerFunc(handlerPing)
	handler.ServeHTTP(rr, req)

	expectedResponse := "PONG"
	responseBody := rr.Body.String()

	// Compare response body
	if responseBody != expectedResponse {
		t.Errorf("For 'ping' endpoint: expecting response %s. Received: %s", expectedResponse, responseBody)
	}
}
