package main

import (
	"context"
	"fmt"
	"github.com/onurbaran/easy-cache/pkg/cache"
	"time"
)

type Person struct {
	Name  string
	Email string
	Age   int
}

func main() {
	config := cache.DefaultConfig()
	config.BaseTTL = 5 * time.Second
	structCache := cache.NewStructCache(config)

	ctx := context.Background()
	user := Person{Name: "Alice", Email: "alice@example.com", Age: 30}
	err := structCache.Set(ctx, "user1", user)
	if err != nil {
		fmt.Println("Error setting struct in cache:", err)
		return
	}

	data, found, _ := structCache.Get(ctx, "user1")
	if found {
		fmt.Printf("Initially found struct in cache: %+v\n", data.(Person))
	} else {
		fmt.Println("Struct not found in cache initially")
	}

	time.Sleep(6 * time.Second)

	data, found, _ = structCache.Get(ctx, "user1")
	if !found {
		fmt.Println("Struct has expired and is no longer in cache")
	} else {
		fmt.Printf("Struct still found in cache after TTL: %+v\n", data.(Person))
	}
}
