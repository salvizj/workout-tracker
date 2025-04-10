package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles(
			"templates/base.html",
			"templates/home.html",
		))

		data := map[string]interface{}{
			"Title": "Home",
		}

		err := tmpl.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
		}
	}
}
