package service

import (
	"github.com/alprnemn/yollapi/internal/domain"
	"github.com/alprnemn/yollapi/internal/repository"
	ps "github.com/alprnemn/yollapi/internal/service/productservice"
	us "github.com/alprnemn/yollapi/internal/service/userservice"
)

type Service struct {
	User    domain.IUserService
	Product domain.IProductService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User: &us.UserService{
			Repo: repo.User,
		},
		Product: &ps.ProductService{
			Repo: repo.Product,
		},
	}
}
