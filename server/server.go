package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"rest-go/db"
	"rest-go/repository"
	"rest-go/websocket"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Config struct {
	Port        string
	JWTSecret   string
	DataBaseUrl string
}

type Server interface {
	Config() *Config
	Hub() *websocket.Hub
}

type Broker struct {
	config *Config
	router *mux.Router
	hub    *websocket.Hub
}

func (b *Broker) Config() *Config {
	return b.config
}

func (b *Broker) Hub() *websocket.Hub {
	return b.hub
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
		hub:    websocket.NewHub(),
	}
	err = nil
	return
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	var Port = b.config.Port
	b.router = mux.NewRouter()
	binder(b, b.router)
	if strings.Index(b.config.Port, ":") != 0 && !strings.Contains(b.config.Port, ":") {
		Port = ":" + b.config.Port
	}
	handler := cors.Default().Handler(b.router)
	repo, err := db.NewPostgresRepository(b.config.DataBaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	go b.hub.Run()
	repository.SetRepository(repo)
	log.Println("Starting server on port", Port)
	if err := http.ListenAndServe(Port, handler); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
