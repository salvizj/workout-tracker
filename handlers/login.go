package handlers

import (
	"html/template"
	"net/http"
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

	}
}
