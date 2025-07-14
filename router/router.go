package router

import (
    "strings"
    "net/http"
    "github.com/lo0kat/kerebro-api/handlers"
)

func NewRouter() http.Handler {
    mux := http.NewServeMux()

    mux.HandleFunc("/api/v1/words", func(w http.ResponseWriter, r *http.Request) {
        path := strings.TrimPrefix(r.URL.Path, "/api/v1/words")
        path = strings.Trim(path, "/")
    
        switch r.Method {
        case http.MethodGet:
            if path == "" {
                handlers.GetWords(w, r)
            } else {
                handlers.GetWord(w, r, path) // path contains the user ID
            }
        case http.MethodDelete:
            handlers.DeleteWord(w, r, path)
        case http.MethodPost:
            handlers.CreateWord(w, r)
        // case http.MethodPatch:
        //     updateUser(w, r, path)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })


    return mux
}
