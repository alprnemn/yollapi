package domain

import (
	"context"
	cmn "github.com/alprnemn/yollapi/common"
	"time"
)

type IUserRepository interface {
	Create(ctx context.Context, user *User) error
	GetAllUsers(ctx context.Context) ([]User, error)
}

type IUserService interface {
	Register(ctx context.Context, user *User) *cmn.ErrorResponse
	GetAllUsers(ctx context.Context) ([]User, error)
}

type IProductRepository interface {
	Create(ctx context.Context, product *Product) error
}

type IProductService interface {
	Add(ctx context.Context, product *Product) error
}

type IRateLimiter interface {
	Allow(ip string) (bool, time.Duration)
}
