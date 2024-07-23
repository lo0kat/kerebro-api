package services

import (
	"log"
	"net/http"

	"github.com/lo0kat/kerebro-api/internal/store"
)

type WordsService struct {
	store store.Store
}

func NewWordsService(s store.Store) *WordsService {
	return &WordsService{store: s}
}

func (s *WordsService) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /words", s.handleCreateWord)
	router.HandleFunc("GET /words", s.handleListWords)
	router.HandleFunc("GET /words/{wordID}", s.handleGetWord)
	router.HandleFunc("DELETE /words/{wordID}", s.handleDeleteWord)
	router.HandleFunc("PATCH /words/{wordID}", s.handlePatchWord)
}

func (s *WordsService) handleCreateWord(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("POST : Word Created"))
	log.Printf("POST /words hit")

}

func (s *WordsService) handleListWords(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET : List words"))
	log.Printf("GET /words hit")
}

func (s *WordsService) handleGetWord(w http.ResponseWriter, r *http.Request) {
	wordID := r.PathValue("wordID")
	w.Write([]byte("GET : Word ID" + wordID))
	log.Printf("POST /words/{wordsID} hit")
}

func (s *WordsService) handleDeleteWord(w http.ResponseWriter, r *http.Request) {
	wordID := r.PathValue("wordID")
	w.Write([]byte("DELETE : Delete Word ID" + wordID))
	log.Printf("DELETE /words/{wordsID} hit")

}

func (s *WordsService) handlePatchWord(w http.ResponseWriter, r *http.Request) {
	wordID := r.PathValue("wordID")
	w.Write([]byte("PATCH : Word ID" + wordID))
	log.Printf("PATCH /words/{wordsID} hit")

}
