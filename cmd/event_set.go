package main

import (
	"context"
	"fmt"
	"github.com/onurbaran/easy-cache/pkg/cache"
	"github.com/onurbaran/easy-cache/pkg/event"
)

type SetEventListener struct {
	cache *cache.Cache
}

func (l *SetEventListener) OnEvent(e event.Event) {
	ctx := context.Background()
	key := e.Data.(string)
	l.cache.Set(ctx, key, "new_value", 1, "default")
	fmt.Println("Set value for key:", key)
}

func main() {
	eventManager := event.NewEventManager()

	config := cache.DefaultConfig()
	config.EventManager = eventManager

	c := cache.NewCache(config)

	listener := &SetEventListener{cache: c}
	eventManager.RegisterListener("setKey", listener)

	ctx := context.Background()
	data, found, _ := c.Get(ctx, "key1")
	if !found {
		fmt.Println("Initially, key1 not found")
	} else {
		fmt.Println("Initially, found key1:", data)
	}

	eventManager.TriggerEvent(event.Event{Name: "setKey", Data: "key1"})

	data, found, _ = c.Get(ctx, "key1")
	if found {
		fmt.Println("After event, found key1:", data)
	} else {
		fmt.Println("After event, key1 not found")
	}
}
