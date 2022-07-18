package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	db "rest-go/db"
	repository "rest-go/repository"

	"github.com/gorilla/mux"
)

type Config struct {
	Port        string
	JWTSecret   string
	DataBaseUrl string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (broker *Broker, err error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("secret is required")
	}
	if config.DataBaseUrl == "" {
		return nil, errors.New("db is required")
	}
	broker = &Broker{
		config: config,
		router: mux.NewRouter(),
	}
	err = nil
	return
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	repo, err := db.NewPostgresRepository(b.config.DataBaseUrl)
	if err != nil{
		log.Fatal(err)
	}
	repository.SetRepository(repo)
	log.Println("Starting server on port", b.Config().Port)
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("listen and serve: ", err)
	}
}
