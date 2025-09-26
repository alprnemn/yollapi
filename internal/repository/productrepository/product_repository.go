package productrepository

import (
	"context"
	"database/sql"
	"github.com/alprnemn/yollapi/internal/domain"
)

type ProductRepository struct {
	Db *sql.DB
}

func (repo *ProductRepository) Create(ctx context.Context, product *domain.Product) error {
	return nil
}
