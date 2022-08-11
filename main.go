package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"rest-go/handlers"
	"rest-go/server"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var templa *template.Template

func main() {
	templa = template.Must(template.ParseGlob("./templates/*.html"))
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATBASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DataBaseUrl: DATBASE_URL,
	})
	if err != nil {
		log.Fatal(err)
	}
	s.Start(BindRoutes)
}
func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler(s, templa)).Methods(http.MethodGet)
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)
}
