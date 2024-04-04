package token

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type Tokens struct {
	redis *redis.Client
}

func NewTokens(redis *redis.Client) *Tokens {
	return &Tokens{redis: redis}
}

func (t *Tokens) Add(token string, jsonData string, timeOut time.Duration) error {
	err := t.redis.Set(context.Background(), token, jsonData, timeOut).Err()
	if err != nil {
		log.Println("Redis 添加Token失败: ", token)
		return err
	}
	log.Println("Redis 写入Token: ", token)
	return nil
}

func (t *Tokens) GetKeyRandom() (string, error) {
	// 获取所有的键
	key, err := t.redis.RandomKey(context.Background()).Result()
	if err != nil {
		return "", err
	}

	return key, nil
}

func (t *Tokens) RemoveKey(key string) error {
	keys, err := t.redis.Keys(context.Background(), "*").Result()
	if err != nil {
		return err
	}
	if len(keys) == 1 {
		return nil
	}
	// 获取所有的键
	err = t.redis.Del(context.Background(), key).Err()
	if err != nil {
		log.Println("Redis 移除Token 失败:")
		return err
	}
	log.Println("Redis 移除Token:", key)
	return nil
}

func (t *Tokens) GetAllKey() ([]string, error) {
	// 获取所有的键
	result, err := t.redis.Keys(context.Background(), "*").Result()
	if err != nil {
		log.Println("Redis 移除Token 失败:")
		return nil, err
	}
	return result, nil
}

func (t *Tokens) GetKeyValueRandom() (string, string, error) {
	key, err := t.redis.RandomKey(context.Background()).Result()
	if err != nil {
		return "", "", err
	}
	value, err := t.redis.Get(context.Background(), key).Result()
	if err != nil {
		return "", "", err
	}
	return key, value, nil
}
