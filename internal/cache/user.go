package cache

import "github.com/redis/go-redis/v9"

type UserRepository struct {
	rdb *redis.Client
}

func (r *UserRepository) Get()    {}
func (r *UserRepository) Set()    {}
func (r *UserRepository) Delete() {}
