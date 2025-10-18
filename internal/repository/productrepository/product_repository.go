package productrepository

import (
	"context"
	"database/sql"

	m "github.com/alprnemn/yollapi/internal/models"
)

type ProductRepository struct {
	Db *sql.DB
}

func (repo *ProductRepository) Create(ctx context.Context, product *m.Product) error {
	return nil
}
