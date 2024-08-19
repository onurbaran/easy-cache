package main

import (
	"context"
	"fmt"
	"github.com/onurbaran/easy-cache/pkg/cache"
	"github.com/onurbaran/easy-cache/pkg/event"
)

func main() {
	eventManager := event.NewEventManager()

	config := cache.DefaultConfig()
	config.EventManager = eventManager

	c := cache.NewCache(config)

	ctx := context.Background()
	c.Set(ctx, "user:123", "John Doe", 1, "default")

	data, found, _ := c.Get(ctx, "user:123")
	if found {
		fmt.Println("Before invalidation, found:", data)
	} else {
		fmt.Println("Before invalidation, not found")
	}

	eventManager.TriggerEvent(event.Event{Name: "invalidateUser", Data: "user:123"})

	data, found, _ = c.Get(ctx, "user:123")
	if found {
		fmt.Println("After invalidation, found:", data)
	} else {
		fmt.Println("After invalidation, not found")
	}
}
