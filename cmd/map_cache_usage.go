package main

import (
	"context"
	"fmt"
	"github.com/onurbaran/easy-cache/pkg/cache"
)

func main() {
	config := cache.DefaultConfig()
	mapCache := cache.NewMapCache(config)

	ctx := context.Background()
	exampleMap := map[string]interface{}{"key1": "value1", "key2": "value2"}
	mapCache.Set(ctx, "example", exampleMap)

	data, found, err := mapCache.Get(ctx, "example")
	if err != nil {
		fmt.Println("Error getting map from cache:", err)
		return
	}
	if found {
		fmt.Printf("Found map in cache: %+v\n", data)
	} else {
		fmt.Println("Map not found in cache")
	}

	err = mapCache.Delete(ctx, "example")
	if err != nil {
		fmt.Println("Error deleting map from cache:", err)
		return
	}

	data, found, err = mapCache.Get(ctx, "example")
	if !found {
		fmt.Println("Map successfully deleted from cache")
	} else {
		fmt.Println("Map still found in cache:", data)
	}
}
