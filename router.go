package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Endpoint struct {
	Path string
	Code uint16
}

func (e Endpoint) Handler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func NewRouter(endpoint Endpoint) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(endpoint.Path, endpoint.Handler())

	return r
}
