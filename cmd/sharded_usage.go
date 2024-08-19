package main

import (
	"context"
	"fmt"
	"github.com/onurbaran/easy-cache/pkg/cache"
	"github.com/onurbaran/easy-cache/pkg/sharded"
	"time"
)

func main() {
	config := cache.DefaultConfig()
	config.BaseTTL = 15 * time.Minute
	config.MaxItems = 200
	config.NumShards = 8

	shardedCache := sharded.NewShardedCache(config)

	ctx := context.Background()
	err := shardedCache.Set(ctx, "key1", "value1", 1, "default")
	if err != nil {
		fmt.Println("Error setting sharded cache:", err)
		return
	}

	data, found, err := shardedCache.Get(ctx, "key1")
	if err != nil {
		fmt.Println("Error getting sharded cache:", err)
		return
	}

	if found {
		fmt.Println("Found in sharded cache:", data)
	} else {
		fmt.Println("Not found in sharded cache")
	}

	err = shardedCache.Delete(ctx, "key1")
	if err != nil {
		fmt.Println("Error deleting from sharded cache:", err)
		return
	}

	data, found, err = shardedCache.Get(ctx, "key1")
	if err != nil {
		fmt.Println("Error getting sharded cache:", err)
		return
	}

	if found {
		fmt.Println("Found in sharded cache:", data)
	} else {
		fmt.Println("Not found in sharded cache after deletion")
	}
}
