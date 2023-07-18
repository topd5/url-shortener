package main

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		// Addr:     "localhost:6379",
		Addr:     "redis:6379",
		Password: "", // пароля нет
		DB:       0,  // использовать дефолтную БД
	})
}

func RedisSet(key string, value string, ttlSeconds ...int) error {
	ttl := time.Second * 0

	if len(ttlSeconds) > 0 {
		ttl = time.Second * time.Duration(ttlSeconds[0])
	}

	err := RedisClient.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}

func RedisGet(key string) (string, error) {
	result, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return result, nil
}
