package main

import (
	"log"
	"net/http"
)

func main() {
	endpoint := Endpoint{
		"/create",
		201,
	}

	r := NewRouter(endpoint)
	srv := &http.Server{
		Handler: r,
		Addr:    ":3131",
	}

	log.Fatal(srv.ListenAndServe())
}
