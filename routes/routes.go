package routes

import (
	"net/http"
	"workout_tracker/handlers"
	"workout_tracker/middleware"
)

func RegisterRoutes() *http.ServeMux {
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/login", handlers.LoginHandler)
	router.HandleFunc("/signUp", handlers.SignupHandler)

	router.Handle("/dashboard", middleware.AuthMiddleware(http.HandlerFunc(handlers.DashboardHandler)))

	return router
}
