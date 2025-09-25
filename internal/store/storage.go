package store

import (
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("record not found")
	ErrConflict          = errors.New("resource already exists")
	ErrDuplicateUsername = errors.New("username already exists")
	ErrDuplicateEmail    = errors.New("email already exists")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Users IUserStore
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Users: &UserStore{
			db: db,
		},
	}
}
