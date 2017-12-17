package main

import (
	"github.com/erikperttu/kubernetes-ready-service-tutorial/handlers"
	"github.com/erikperttu/kubernetes-ready-service-tutorial/version"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Printf(
		"Starting the service...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	r := handlers.Router(version.BuildTime, version.Commit, version.Release)
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
