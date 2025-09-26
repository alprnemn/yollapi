package productservice

import (
	"context"
	"github.com/alprnemn/yollapi/internal/domain"
)

type ProductService struct {
	Repo domain.IProductRepository
}

func (service *ProductService) Add(ctx context.Context, product *domain.Product) error {
	return nil
}
