package main

import (
	"fmt"
	"log"
	"net/http"
	"workout_tracker/config"
	"workout_tracker/db"
	"workout_tracker/global"
	"workout_tracker/routes"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	dbConn, err := db.InitDB(config)
	if err != nil {
		log.Fatal("Error initializing database: ", err)
	}
	global.DB = dbConn
	defer global.DB.Close()

	router := routes.RegisterRoutes()

	server := &http.Server{
		Addr:    config.BindAddress + ":" + config.Port,
		Handler: router,
	}

	startupMsg := fmt.Sprintf(
		"Server running on: http://%s:%s\nEnvironment: %s\nDatabase: %s",
		config.Domain,
		config.Port,
		config.Env,
		config.DatabasePath,
	)
	fmt.Println(startupMsg)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Error starting server: ", err)
	}
}
