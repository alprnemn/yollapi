package productservice

import (
	"context"

	"github.com/alprnemn/yollapi/internal/models"
)

type ProductService struct {
	Repo models.IProductRepository
}

func (service *ProductService) Add(ctx context.Context, product *models.Product) error {
	return nil
}
