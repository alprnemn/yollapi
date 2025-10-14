package userservice

import (
	"context"

	cmn "github.com/alprnemn/yollapi/common"
	"github.com/alprnemn/yollapi/internal/domain"
)

type UserService struct {
	Repo domain.IUserRepository
}

func (service *UserService) Register(ctx context.Context, user *domain.User) *cmn.ErrorResponse {

	if err := service.Repo.Create(ctx, user); err != nil {
		err := handleErrorRegister(err)
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *UserService) GetAllUsers(ctx context.Context) ([]domain.User, error) {

	users, err := service.Repo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
