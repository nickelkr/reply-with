package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEndpointPath(t *testing.T) {
	endpoint := Endpoint{
		"/test",
		200,
	}

	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	recr := httptest.NewRecorder()
	router := NewRouter(endpoint)
	router.ServeHTTP(recr, req)

	// TODO: test body instead of Code
	if status := recr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v expected %v", status, http.StatusOK)
	}
}

func TestEndpointCode(t *testing.T) {
	endpoint := Endpoint{
		"/test",
		503,
	}

	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	recr := httptest.NewRecorder()
	router := NewRouter(endpoint)
	router.ServeHTTP(recr, req)

	if status := recr.Code; status != 503 {
		t.Errorf("Wrong status code: go %v expected %v", status, 503)
	}
}
