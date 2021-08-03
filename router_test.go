package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSimpleGet(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	recr := httptest.NewRecorder()

	router := NewRouter()
	router.ServeHTTP(recr, req)

	if status := recr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: go %v need %v", status, http.StatusOK)
	}
}
