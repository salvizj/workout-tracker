package handlers

import (
	"html/template"
	"net/http"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/signup.html",
	))

	data := map[string]interface{}{
		"Title": "Sign Up",
	}

	err := tmpl.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		http.Error(w, "Failed to render signup page", http.StatusInternalServerError)
	}
}
