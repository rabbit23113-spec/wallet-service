package main

import (
	"context"
	"log/slog"
	"net/http"
	"wallet/internal/app/api"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(handler http.Handler) *Server {
	httpServer := &http.Server{
		Addr:           ":8080",
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
	handler := api.NewHandler()
	srv := NewServer(handler)
	srv.Start()
}
