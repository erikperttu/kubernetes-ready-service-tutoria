package main

import (
	"github.com/erikperttu/kubernetes-ready-service-tutorial/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("Starting the service...")
	router := handlers.Router()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8000", router))
}
