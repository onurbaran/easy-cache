package main

import (
	"context"
	"fmt"
	"github.com/onurbaran/easy-cache/pkg/cache"
)

func main() {
	config := cache.DefaultConfig()
	sliceCache := cache.NewSliceCache(config)

	ctx := context.Background()
	exampleSlice := []interface{}{"value1", "value2", "value3"}
	sliceCache.Set(ctx, "exampleSlice", exampleSlice)

	data, found, err := sliceCache.Get(ctx, "exampleSlice")
	if err != nil {
		fmt.Println("Error getting slice from cache:", err)
		return
	}
	if found {
		fmt.Printf("Found slice in cache: %+v\n", data)
	} else {
		fmt.Println("Slice not found in cache")
	}

	err = sliceCache.Delete(ctx, "exampleSlice")
	if err != nil {
		fmt.Println("Error deleting slice from cache:", err)
		return
	}

	data, found, err = sliceCache.Get(ctx, "exampleSlice")
	if !found {
		fmt.Println("Slice successfully deleted from cache")
	} else {
		fmt.Println("Slice still found in cache:", data)
	}
}
