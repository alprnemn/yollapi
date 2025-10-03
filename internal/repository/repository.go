package repository

import (
	"database/sql"

	"github.com/alprnemn/yollapi/internal/domain"
	pr "github.com/alprnemn/yollapi/internal/repository/productrepository"
	ur "github.com/alprnemn/yollapi/internal/repository/userrepository"
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
