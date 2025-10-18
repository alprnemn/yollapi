package userservice

import (
	"context"

	cmn "github.com/alprnemn/yollapi/common"
	m "github.com/alprnemn/yollapi/internal/models"
)

type UserService struct {
	Repo m.IUserRepository
}

func (service *UserService) Register(ctx context.Context, user *m.User) *cmn.ErrorResponse {

	if err := service.Repo.Create(ctx, user); err != nil {
		err := handleErrorRegister(err)
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *UserService) GetAllUsers(ctx context.Context) ([]m.User, error) {

	users, err := service.Repo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
