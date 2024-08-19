package cache

import (
	"context"
)

type MapCache struct {
	cache *Cache
}

func NewMapCache(config *Config) *MapCache {
	return &MapCache{
		cache: NewCache(config),
	}
}

func (mc *MapCache) Set(ctx context.Context, key string, value map[string]interface{}) error {
	return mc.cache.Set(ctx, key, value, 1, "map")
}

func (mc *MapCache) Get(ctx context.Context, key string) (map[string]interface{}, bool, error) {
	data, found, err := mc.cache.Get(ctx, key)
	if !found || err != nil {
		return nil, false, err
	}
	return data.(map[string]interface{}), true, nil
}

func (mc *MapCache) Delete(ctx context.Context, key string) error {
	return mc.cache.Delete(ctx, key)
}
