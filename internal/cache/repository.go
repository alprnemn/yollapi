package cache

import (
	"github.com/alprnemn/yollapi/internal/models"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	User models.ICacheUserRepository
}

func NewRedisRepository(rdb *redis.Client) RedisRepository {
	return RedisRepository{
		User: &UserRepository{rdb: rdb},
	}
}
