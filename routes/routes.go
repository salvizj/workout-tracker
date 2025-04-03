package routes

import (
	"net/http"
	"workout_tracker/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/", handlers.HomeHandler)
}
