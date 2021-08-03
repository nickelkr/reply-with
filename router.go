package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthHandler)

	return r
}
