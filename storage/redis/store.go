package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type CacheItem struct {
	Key     string
	Value   []byte
	Headers map[string][]string
}

type CacheStore struct {
	client *redis.Client
}

func NewCacheStore(client *redis.Client) *CacheStore {
	return &CacheStore{
		client,
	}
}

func (s *CacheStore) Exists(key string) bool {
	ctx := context.TODO()
	res, err := s.client.Exists(ctx, key).Result()
	if err != nil {
		return false
	}
	return res == 1
}

func (s *CacheStore) Save(key string, data []byte, headers map[string][]string) error {
	ctx := context.TODO()

	jsonData, err := json.Marshal(CacheItem{
		Key:     key,
		Value:   data,
		Headers: headers,
	})
	if err != nil {
		return fmt.Errorf("error marshaling data: %w", err)
	}

	err = s.client.Set(ctx, key, jsonData, 0).Err()
	if err != nil {
		return fmt.Errorf("unable to save item in redis with key '%s': %w", key, err)
	}
	return nil
}

func (s *CacheStore) Get(key string) ([]byte, map[string][]string, error) {
	ctx := context.TODO()
	val, err := s.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, nil, fmt.Errorf("unable to get item: %w", err)
	}
	var item CacheItem
	err = json.Unmarshal(val, &item)
	if err != nil {
		return nil, nil, fmt.Errorf("error unmarshaling json: %w", err)
	}
	return item.Value, item.Headers, nil
}

func (s *CacheStore) DeleteById(id string) error {
	ctx := context.TODO()
	err := s.client.Del(ctx, id).Err()
	if err != nil {
		return fmt.Errorf("redis delete error: %w", err)
	}
	return nil
}
