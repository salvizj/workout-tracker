package main

import (
	"log"
	"net/http"
	"workout_tracker/config"
	"workout_tracker/routes"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	routes.RegisterRoutes()

	err = http.ListenAndServe(":"+cfg.PORT, nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
