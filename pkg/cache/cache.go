package cache

import (
	"context"
	"github.com/onurbaran/easy-cache/pkg/event"
	"sync"
)

type Cache struct {
	items  map[string]*Item
	mu     sync.RWMutex
	config *Config
}

func NewCache(config *Config) *Cache {
	if config == nil {
		config = DefaultConfig()
	}

	return &Cache{
		items:  make(map[string]*Item),
		config: config,
	}
}

func (c *Cache) Set(ctx context.Context, key string, data interface{}, priority int, category string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		c.mu.Lock()
		if len(c.items) >= c.config.MaxItems {
			c.mu.Unlock()
			c.config.EvictionPolicy.Evict(c)
			c.mu.Lock()
		}

		ttl := c.config.BaseTTL
		if overrideTTL, found := c.config.TTLOverrides[category]; found {
			ttl = overrideTTL
		}

		item := NewItem(key, data, priority, ttl)
		c.items[key] = item
		c.mu.Unlock()

		if c.config.EventManager != nil {
			c.config.EventManager.TriggerEvent(event.Event{Name: "itemAdded", Data: key})
		}

		return nil
	}
}

func (c *Cache) Get(ctx context.Context, key string) (interface{}, bool, error) {
	select {
	case <-ctx.Done():
		return nil, false, ctx.Err()
	default:
		c.mu.RLock()
		defer c.mu.RUnlock()

		item, found := c.items[key]
		if !found || item.IsExpired() {
			return nil, false, nil
		}

		item.IncrementAccessCount().RefreshTTL()

		if c.config.Serializer != nil {
			var deserializedData interface{}
			err := c.config.Serializer.Deserialize([]byte(item.GetData().(string)), &deserializedData)
			if err == nil {
				return deserializedData, true, nil
			}
			return nil, false, err
		}

		return item.GetData(), true, nil
	}
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		c.mu.Lock()
		defer c.mu.Unlock()

		delete(c.items, key)

		if c.config.EventManager != nil {
			c.config.EventManager.TriggerEvent(event.Event{Name: "itemDeleted", Data: key})
		}

		return nil
	}
}

func (c *Cache) Items() map[string]*Item {
	c.mu.RLock()
	defer c.mu.RUnlock()

	itemsCopy := make(map[string]*Item)
	for k, v := range c.items {
		itemsCopy[k] = v
	}
	return itemsCopy
}
