package redis

import (
	"github.com/redis/go-redis/v9"
)

func NewConnection(c Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Username: c.Username,
		Password: c.Password,
		DB:       c.DB,
	})
	return rdb
}
