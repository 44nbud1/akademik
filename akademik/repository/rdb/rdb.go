package rdb

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"log"
	"time"
)

type redisRepo struct {
	redis *redis.Client
}

func NewRedisRepo(redis *redis.Client) *redisRepo {
	return &redisRepo{redis: redis}
}

func (r *redisRepo) SetDataWithExpiry(key, value string, expiredPeriod time.Duration) error {

	if err := r.redis.Set(key, value, 0).Err(); err != nil {
		fmt.Println(err)
	}

	return nil
}

func (r *redisRepo) GetData(key string) (string, error) {
	val, err := r.redis.Get(key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *redisRepo) FlushData() error {
	log.Print("Clear redis data")
	return r.redis.FlushAll().Err()
}

func (r *redisRepo) Ping() {
	fmt.Println("Redis Ping: -> Response : ", r.redis.Ping())
}
