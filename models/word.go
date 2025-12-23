package models

import (
	"time"
)

type Word struct {
	ID           int       `json:"id"`
	Kr           string    `json:"kr_word"`
	En           string    `json:"en_word"`
	CreationDate time.Time `json:"created_at"`
}
