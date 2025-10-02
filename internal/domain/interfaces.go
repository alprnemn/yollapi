package domain

import (
	"context"
	cmn "github.com/alprnemn/yollapi/common"
)

type IUserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type IUserService interface {
	Register(ctx context.Context, user *User) *cmn.ErrorResponse
}

type IProductRepository interface {
	Create(ctx context.Context, product *Product) error
}

type IProductService interface {
	Add(ctx context.Context, product *Product) error
}
