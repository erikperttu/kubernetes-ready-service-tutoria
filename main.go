package main

import (
	"net/http"
	"log"
	"github.com/erikperttu/kubernetes-ready-service-tutoral/handlers"
	"os"
)

func main(){
	log.Print("Starting the service...")
	router := handlers.Router()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8000", router))
}


