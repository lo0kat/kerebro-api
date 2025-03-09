package store

import(
	"database/sql"
	)

type Store interface {
	CreateInstance() error
}

// type Word struct {
// 	en_word: string, `json: en_word`
// 	kr_word: string, `json: kr_word`
// 	created_at: time.Time, `json: created_at`
// }

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateInstance() error {
	return nil
}
