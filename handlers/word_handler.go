package handlers

import(
	"encoding/json"
	"net/http"
	"github.com/lo0kat/kerebro-api/models"
	"github.com/lo0kat/kerebro-api/repository"
)

func GetWords(w http.ResponseWriter, r *http.Request) {
    words, err := repository.GetAllWords()
    if err != nil {
        http.Error(w, "Failed to fetch words", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(words)
}

func CreateWord(w http.ResponseWriter, r *http.Request) {
    var word models.Word
    if err := json.NewDecoder(r.Body).Decode(&word); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

	id, err := repository.CreateWord(word)
    if err != nil {
        http.Error(w, "Failed to create word", http.StatusInternalServerError)
        return
    }

	resp := map[string]string{"word_id": id}
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func GetWord(w http.ResponseWriter, r *http.Request, id string) {
    word, err := repository.GetWord(id)
    if err != nil {
        http.Error(w, "Failed to fetch word", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(word)
}

func DeleteWord(w http.ResponseWriter, r *http.Request, id string) {
	if err := repository.DeleteWord(id); err != nil {
		http.Error(w, "Failed to delete word", http.StatusInternalServerError)
		return
	}
    w.Header().Set("Content-Type", "application/json")
}

func PatchWord(w http.ResponseWriter, r *http.Request, id string) {
	if err := repository.DeleteWord(id); err != nil {
		http.Error(w, "Failed to delete word", http.StatusInternalServerError)
		return
	}
    w.Header().Set("Content-Type", "application/json")
}