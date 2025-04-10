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
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}
	global.CONFIG = cfg

	dbConn, err := db.InitDB(global.CONFIG)
	if err != nil {
		log.Fatal("Error initializing database: ", err)
	}
	global.DB = dbConn
	defer global.DB.Close()

	router := routes.RegisterRoutes()

	server := &http.Server{
		Addr:    global.CONFIG.BindAddress + ":" + global.CONFIG.Port,
		Handler: router,
	}

	startupMsg := fmt.Sprintf(
		"Server running on: http://%s:%s\nEnvironment: %s\nDatabase: %s",
		global.CONFIG.Domain,
		global.CONFIG.Port,
		global.CONFIG.Env,
		global.CONFIG.DatabasePath,
	)
	fmt.Println(startupMsg)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Error starting server: ", err)
	}
}
