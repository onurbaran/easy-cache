package sharded

import (
	"context"
	"github.com/onurbaran/easy-cache/pkg/cache"
	"hash/fnv"
)

type ShardedCache struct {
	shards    []*cache.Cache
	numShards int
}

func NewShardedCache(config *cache.Config) *ShardedCache {
	shards := make([]*cache.Cache, config.NumShards)
	for i := 0; i < config.NumShards; i++ {
		shards[i] = cache.NewCache(config)
	}
	return &ShardedCache{
		shards:    shards,
		numShards: config.NumShards,
	}
}

func (sc *ShardedCache) getShard(key string) *cache.Cache {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	shardIndex := int(hash.Sum32()) % sc.numShards
	return sc.shards[shardIndex]
}

func (sc *ShardedCache) Set(ctx context.Context, key string, data interface{}, priority int, category string) error {
	shard := sc.getShard(key)
	return shard.Set(ctx, key, data, priority, category)
}

func (sc *ShardedCache) Get(ctx context.Context, key string) (interface{}, bool, error) {
	shard := sc.getShard(key)
	return shard.Get(ctx, key)
}

func (sc *ShardedCache) Delete(ctx context.Context, key string) error {
	shard := sc.getShard(key)
	return shard.Delete(ctx, key)
}
