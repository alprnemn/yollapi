package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	c "github.com/alprnemn/yollapi/common"
	m "github.com/alprnemn/yollapi/internal/models"
	"github.com/redis/go-redis/v9"
)

type UserCacheRepository struct {
	rdb *redis.Client
}

func (r *UserCacheRepository) Get(ctx context.Context, userID int64) (*m.User, error) {
	cacheKey := fmt.Sprintf("user-%d", userID)

	data, err := r.rdb.Get(ctx, cacheKey).Result()
	if errors.Is(err, redis.Nil) {
		return nil, c.ErrNotFoundOnRedis
	} else if err != nil {
		return nil, err
	}

	var user m.User

	if data != "" {
		err := json.Unmarshal([]byte(data), &user)
		fmt.Printf("data on redis :\n--id: %d \n--username: %s, \n--email: %s\n", user.ID, user.Username, user.Email)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil

}

func (r *UserCacheRepository) Set(ctx context.Context, user *m.User) error {
	cacheKey := fmt.Sprintf("user-%d", user.ID)

	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	if err := r.rdb.SetEx(ctx, cacheKey, data, c.UserExpTime).Err(); err != nil {
		return c.ErrSetUserCache
	}
	fmt.Println("setted data: ", string(data))
	return nil
}

func (r *UserCacheRepository) Delete(ctx context.Context, userID int64) {
	cacheKey := fmt.Sprintf("user-%d", userID)
	r.rdb.Del(ctx, cacheKey)
}
