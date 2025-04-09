package handlers

import (
	"html/template"
	"net/http"
	"workout_tracker/services"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/signUp.html"))
		data := map[string]interface{}{
			"Title": "Sign Up",
		}
		err := tmpl.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			http.Error(w, "Failed to render signup page", http.StatusInternalServerError)
			return
		}
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Invalid form submission", http.StatusBadRequest)
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		err = services.SignUp(email, password)
		if err != nil {
			http.Error(w, "Sign up failed: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
