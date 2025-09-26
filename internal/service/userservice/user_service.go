package userservice

import (
	"context"
	"github.com/alprnemn/yollapi/internal/domain"
	"log"
)

type UserService struct {
	Repo domain.IUserRepository
}

func (service *UserService) Register(ctx context.Context, user *domain.User) error {
	log.Print("from user service")

	if err := service.Repo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}
