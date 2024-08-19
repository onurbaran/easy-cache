package main

import (
	"context"
	"fmt"
	"github.com/onurbaran/easy-cache/pkg/cache"
)

func main() {
	config := cache.DefaultConfig()
	config.EvictionPolicy = &cache.LFUEviction{}
	config.MaxItems = 2

	c := cache.NewCache(config)

	ctx := context.Background()
	c.Set(ctx, "key1", "value1", 1, "default")
	c.Set(ctx, "key2", "value2", 1, "default")

	c.Get(ctx, "key1")
	c.Get(ctx, "key1")

	c.Set(ctx, "key3", "value3", 1, "default")

	if _, found, _ := c.Get(ctx, "key1"); found {
		fmt.Println("Found key1 (should not be evicted)")
	} else {
		fmt.Println("Key1 evicted")
	}

	if _, found, _ := c.Get(ctx, "key2"); found {
		fmt.Println("Found key2 (should be evicted)")
	} else {
		fmt.Println("Key2 evicted")
	}

	if _, found, _ := c.Get(ctx, "key3"); found {
		fmt.Println("Found key3 (should not be evicted)")
	} else {
		fmt.Println("Key3 evicted")
	}
}
