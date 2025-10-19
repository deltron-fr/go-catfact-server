package main

import (
	"log"
	"net/http"
	"time"

	"github.com/deltron-fr/hng-go/internal/handlers"
)

const port = ":8080"

func main() {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr: port,
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	mux.HandleFunc("/me", handlers.HandleMe)
	
	log.Printf("Server running on %s", port)
	log.Fatal(server.ListenAndServe())

}