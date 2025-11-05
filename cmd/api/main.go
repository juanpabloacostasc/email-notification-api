package main

import (
	"log"
	"net/http"

	"notification-service/cmd/api/server"
)

func main() {
	server.MapURLs()

	log.Println("Server starting on port 8082...")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
