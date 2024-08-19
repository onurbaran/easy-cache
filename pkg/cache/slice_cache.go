package cache

import (
	"context"
)

type SliceCache struct {
	cache *Cache
}

func NewSliceCache(config *Config) *SliceCache {
	return &SliceCache{
		cache: NewCache(config),
	}
}

func (sc *SliceCache) Set(ctx context.Context, key string, value []interface{}) error {
	return sc.cache.Set(ctx, key, value, 1, "slice")
}

func (sc *SliceCache) Get(ctx context.Context, key string) ([]interface{}, bool, error) {
	data, found, err := sc.cache.Get(ctx, key)
	if !found || err != nil {
		return nil, false, err
	}
	return data.([]interface{}), true, nil
}

func (sc *SliceCache) Delete(ctx context.Context, key string) error {
	return sc.cache.Delete(ctx, key)
}
