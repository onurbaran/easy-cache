package main

import (
	"context"
	"fmt"
	"github.com/onurbaran/easy-cache/pkg/cache"
	"github.com/onurbaran/easy-cache/pkg/event"
)

type User struct {
	Name  string
	Email string
	Age   int
}

type InvalidateEventListener struct {
	cache *cache.StructCache
}

func (l *InvalidateEventListener) OnEvent(e event.Event) {
	ctx := context.Background()
	key := e.Data.(string)
	l.cache.Delete(ctx, key)
	fmt.Println("Invalidated struct with key:", key)
}

func main() {
	eventManager := event.NewEventManager()

	config := cache.DefaultConfig()
	config.EventManager = eventManager
	structCache := cache.NewStructCache(config)

	ctx := context.Background()
	user := User{Name: "Alice", Email: "alice@example.com", Age: 30}
	structCache.Set(ctx, "user1", user)

	data, found, _ := structCache.Get(ctx, "user1")
	if found {
		fmt.Printf("Found struct in cache: %+v\n", data.(User))
	} else {
		fmt.Println("Struct not found in cache")
	}

	listener := &InvalidateEventListener{cache: structCache}
	eventManager.RegisterListener("invalidateUser", listener)

	eventManager.TriggerEvent(event.Event{Name: "invalidateUser", Data: "user1"})

	data, found, _ = structCache.Get(ctx, "user1")
	if !found {
		fmt.Println("Struct successfully invalidated and deleted from cache")
	} else {
		fmt.Println("Struct still found in cache:", data)
	}
}
