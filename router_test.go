package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEndpointPath(t *testing.T) {
	endpoint := Endpoint{
		Path: "/test",
		Code: 200,
	}

	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	recr := httptest.NewRecorder()
	router := NewRouter(endpoint)
	router.ServeHTTP(recr, req)

	if status := recr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v expected %v", status, http.StatusOK)
	}
}

func TestEndpointCode(t *testing.T) {
	endpoint := Endpoint{
		Path: "/test",
		Code: 503,
	}

	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	recr := httptest.NewRecorder()
	router := NewRouter(endpoint)
	router.ServeHTTP(recr, req)

	if status := recr.Code; status != 503 {
		t.Errorf("Wrong status code: got %v expected %v", status, 503)
	}
}

func TestEndpointBody(t *testing.T) {
	payload := []byte("some data")
	endpoint := Endpoint{
		"/with_body",
		200,
		payload,
	}

	req, err := http.NewRequest("GET", "/with_body", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	router := NewRouter(endpoint)
	router.ServeHTTP(rec, req)

	if body := rec.Body.Bytes(); !bytes.Equal(body, payload) {
		t.Errorf("Incorrect data: got %v expected %v", body, payload)
	}
}
