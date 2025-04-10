package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles(
			"templates/base.html",
			"templates/dashboard.html",
		))

		data := map[string]interface{}{
			"Title": "Dashboard",
		}

		err := tmpl.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
		}
	}
}
