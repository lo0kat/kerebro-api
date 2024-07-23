package api

import (
	"log"
	"net/http"

	"github.com/lo0kat/kerebro-api/internal/services"
	"github.com/lo0kat/kerebro-api/internal/store"
)

type APIServer struct {
	addr  string
	store store.Store
}

func NewAPIServer(addr string, store store.Store) *APIServer {
	return &APIServer{
		addr:  addr,
		store: store,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	baseRouter := http.NewServeMux()
	wordService := services.NewWordsService(nil)

	wordService.RegisterRoutes(router)
	baseRouter.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	server := http.Server{
		Addr:    s.addr,
		Handler: baseRouter,
	}

	log.Printf("Server has started on %s", s.addr)

	return server.ListenAndServe()
}
