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
		return handleErrorRegister(err)
	}
	return nil
}
