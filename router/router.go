package router

import (
	"net/http"
	"strings"

	"github.com/lo0kat/kerebro-api/handlers"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/words", wordsHandler)
	mux.HandleFunc("/api/v1/words/", wordsHandler)

	return mux
}

func wordsHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/v1/words")
	path = strings.Trim(path, "/")

	hasID := path != ""

	switch r.Method {

	case http.MethodGet:
		if hasID {
			handlers.GetWord(w, r, path)
		} else {
			handlers.GetWords(w, r)
		}

	case http.MethodPost:
		if hasID {
			http.Error(w, "POST not allowed on resource", http.StatusMethodNotAllowed)
			return
		}
		handlers.CreateWord(w, r)

	case http.MethodDelete:
		if !hasID {
			http.Error(w, "DELETE requires resource ID", http.StatusBadRequest)
			return
		}
		handlers.DeleteWord(w, r, path)

	// case http.MethodPatch:
	// 	if !hasID {
	// 		http.Error(w, "PATCH requires resource ID", http.StatusBadRequest)
	// 		return
	// 	}
	// 	handlers.UpdateWord(w, r, path)

	// case http.MethodPut:
	// 	if !hasID {
	// 		http.Error(w, "PUT requires resource ID", http.StatusBadRequest)
	// 		return
	// 	}
	// 	handlers.UpdateWord(w, r, path)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
