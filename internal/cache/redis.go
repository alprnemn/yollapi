package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

func NewRedisClient(addr, pw string, db int) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pw,
		DB:       db,
	})

	ctx := context.Background()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("error connecting redis: ", err)
	}

	log.Println("redis connected successfully")

	return rdb
}
