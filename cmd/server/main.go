package main

import (
	"context"
	"log/slog"
	"net/http"
	"wallet/internal/app/api"
	config2 "wallet/internal/app/config"
	"wallet/internal/app/repository"
	service2 "wallet/internal/app/service"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(config *config2.Config, handler http.Handler) *Server {
	httpServer := &http.Server{
		Addr:           config.Port,
		MaxHeaderBytes: 1 << 20,
		Handler:        handler,
	}
	return &Server{httpServer: httpServer}
}

func (server *Server) Start() {
	err := server.httpServer.ListenAndServe()
	if err != nil {
		slog.Error(err.Error())
		return
	}
}

func (server *Server) Shutdown() {
	err := server.httpServer.Shutdown(context.Background())
	if err != nil {
		slog.Error(err.Error())
	}
}

func main() {
	config := config2.NewConfig()
	repo := repository.New(config)
	service := service2.NewService(repo)
	handler := api.NewHandler(service)
	srv := NewServer(config, handler)
	srv.Start()
}
