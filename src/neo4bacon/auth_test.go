package neo4bacon

import (
	"reflect"
	"testing"
)

func TestNewAuth(t *testing.T) {
	observed := newAuth()
	expected := &Auth{}
	if !reflect.DeepEqual(observed, expected) {
		t.Errorf("Expected: %v. Observed: %v", expected, observed)
	}
}

func TestGetURL(t *testing.T) {
	a := &Auth{
		user:     "user",
		password: "password",
		baseURL:  "baseURL",
		port:     "port",
	}
	a.getURL()
	expected := "bolt://user:password@baseURL:port/"
	if a.URL != expected {
		t.Errorf("Expected %s. Observed: %s", expected, a.URL)
	}
}
