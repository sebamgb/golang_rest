package handlers

import (
	"html/template"
	"net/http"
	"rest-go/models"
	"rest-go/server"
)

type CR struct {
	Response models.ClientRespose
}

type HomeResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func HomeHandler(s server.Server, templa *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		data := CR{
			Response: models.ClientRespose{
				Title:   "Bienvenido",
				Actions: "Puedes registrarte",
				// UserActive: true,
			},
		}
		templa.ExecuteTemplate(w, "index.html", data)
	}
}
