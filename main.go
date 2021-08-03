package main

import (
	"log"
	"net/http"
)

func main() {
	r := NewRouter()
	srv := &http.Server{
		Handler: r,
		Addr:    ":3131",
	}

	log.Fatal(srv.ListenAndServe())
}
