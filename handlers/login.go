package handlers

import (
	"html/template"
	"net/http"
	"time"
	"workout_tracker/services"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles(
			"templates/base.html",
			"templates/login.html",
		))

		data := map[string]interface{}{
			"Title": "Login",
		}

		err := tmpl.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			http.Error(w, "Failed to render login page", http.StatusInternalServerError)
		}
	}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Invalid form submission", http.StatusBadRequest)
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		token, err := services.Login(email, password)
		if err != nil {
			http.Error(w, "Login failed: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "auth_token",
			Value:    token,
			Path:     "/",
			HttpOnly: false,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
			Expires:  time.Now().Add(24 * time.Hour), // Match JWT expiry
		})

		http.Redirect(w, r, "/", http.StatusFound)
	}
}
