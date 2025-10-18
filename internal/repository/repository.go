package repository

import (
	"database/sql"

	"github.com/alprnemn/yollapi/internal/models"
	pr "github.com/alprnemn/yollapi/internal/repository/productrepository"
	ur "github.com/alprnemn/yollapi/internal/repository/userrepository"
)

type Repository struct {
	User    models.IUserRepository
	Product models.IProductRepository
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
