package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	port := ":8080"
	server := &http.Server{
		Addr:         port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Printf("Listening at %s", port)
	log.Fatal(server.ListenAndServe())
}
