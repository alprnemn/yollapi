package domain

import "context"

type IUserRepository interface {
	Create(ctx context.Context, user *User) error
}

type IUserService interface {
	Register(ctx context.Context, user *User) error
}

type IProductRepository interface {
	Create(ctx context.Context, product *Product) error
}

type IProductService interface {
	Add(ctx context.Context, product *Product) error
}
