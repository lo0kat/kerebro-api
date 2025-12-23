package repository

import (
	"context"
	"log"
	"time"

	"github.com/lo0kat/kerebro-api/db"
	"github.com/lo0kat/kerebro-api/models"
)

func GetAllWords() ([]models.Word, error) {
	rows, err := db.Pool.Query(context.Background(), "select * from Words;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	words := []models.Word{}
	for rows.Next() {
		var word models.Word
		if err := rows.Scan(&word.ID, &word.Kr, &word.En, &word.CreationDate); err != nil {
			return nil, err
		}
		words = append(words, word)
	}

	return words, nil
}

func CreateWord(word models.Word) (string, error) {
	var id string
	err := db.Pool.QueryRow(context.Background(), "INSERT INTO Words (kr_word, en_word, created_at) VALUES ($1, $2, $3) RETURNING id;", &word.Kr, &word.En, &word.CreationDate).Scan(&id)
	if err != nil {
		log.Fatalf("Insert failed: %v", err)
		return "", err
	}
	return id, nil
}

func GetWord(id string) (*models.Word, error) {
	var word models.Word
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := db.Pool.QueryRow(ctx, "select * from Words where id=$1;", id).Scan(&word.ID, &word.Kr, &word.En, &word.CreationDate); err != nil {
		return nil, err
	}
	return &word, nil
}

func DeleteWord(id string) error {
	_, err := db.Pool.Exec(context.Background(), "Delete from Words where id=$1;", id)
	return err
}

func UpdateWord(id string, updatedWord *models.Word) error {
	_, err := db.Pool.Exec(context.Background(), "Update Words SET kr_word=$1 en_word=$2 where id=$3;", updatedWord.Kr, updatedWord.En, id)
	return err
}
