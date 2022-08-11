package handlers

import (
	"html/template"
	"net/http"
	"rest-go/server"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(s server.Server, templa *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		data := HomeResponse{
			Message: "Bienvenido",
			Status:  true,
		}
		templa.ExecuteTemplate(w, "index.html", data)
	}
}
