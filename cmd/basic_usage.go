package main

import (
	"context"
	"fmt"
	"github.com/onurbaran/easy-cache/pkg/cache"
	"time"
)

func main() {
	config := cache.DefaultConfig()
	config.BaseTTL = 10 * time.Minute
	config.MaxItems = 100

	config.EvictionPolicy = &cache.LRUEviction{}

	c := cache.NewCache(config)

	ctx := context.Background()
	err := c.Set(ctx, "key1", "value1", 1, "default")
	if err != nil {
		fmt.Println("Error setting cache:", err)
		return
	}

	data, found, err := c.Get(ctx, "key1")
	if err != nil {
		fmt.Println("Error getting cache:", err)
		return
	}

	if found {
		fmt.Println("Found:", data)
	} else {
		fmt.Println("Not found")
	}

	err = c.Delete(ctx, "key1")
	if err != nil {
		fmt.Println("Error deleting cache:", err)
		return
	}

	data, found, err = c.Get(ctx, "key1")
	if err != nil {
		fmt.Println("Error getting cache:", err)
		return
	}

	if found {
		fmt.Println("Found:", data)
	} else {
		fmt.Println("Not found after deletion")
	}
}
