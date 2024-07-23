package store

import "database/sql"

type Store interface {
	CreateInstance() error
}

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
