package repository

import (
	"database/sql"
	"errors"
	"github.com/alprnemn/yollapi/internal/domain"
	pr "github.com/alprnemn/yollapi/internal/repository/productrepository"
	ur "github.com/alprnemn/yollapi/internal/repository/userrepository"
	"time"
)

var (
	ErrNotFound          = errors.New("record not found")
	ErrConflict          = errors.New("resource already exists")
	ErrDuplicateUsername = errors.New("username already exists")
	ErrDuplicateEmail    = errors.New("email already exists")
	QueryTimeoutDuration = time.Second * 5
)

type Repository struct {
	User    domain.IUserRepository
	Product domain.IProductRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User: &ur.UserRepository{
			Db: db,
		},
		Product: &pr.ProductRepository{
			Db: db,
		},
	}
}
