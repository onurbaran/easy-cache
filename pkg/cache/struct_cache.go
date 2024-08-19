package cache

import (
	"context"
)

type StructCache struct {
	cache *Cache
}

func NewStructCache(config *Config) *StructCache {
	return &StructCache{
		cache: NewCache(config),
	}
}

func (sc *StructCache) Set(ctx context.Context, key string, value interface{}) error {
	return sc.cache.Set(ctx, key, value, 1, "struct")
}

func (sc *StructCache) Get(ctx context.Context, key string) (interface{}, bool, error) {
	return sc.cache.Get(ctx, key)
}

func (sc *StructCache) Delete(ctx context.Context, key string) error {
	return sc.cache.Delete(ctx, key)
}
