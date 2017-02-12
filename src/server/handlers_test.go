package server

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerNotFound(t *testing.T) {

	req, err := http.NewRequest("Get", "/foo", nil)
	if err != nil {
		t.Error("Failed to prepare request: ", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlerNotFound)
	handler.ServeHTTP(rr, req)

	expectedResponseCode := 404
	responseCode := rr.Code

	if expectedResponseCode != rr.Code {
		t.Errorf("For '/foo' endpoint: expecting response code %d. Received: %d", expectedResponseCode, responseCode)
	}
}

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
