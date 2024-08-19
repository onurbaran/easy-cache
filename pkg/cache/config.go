package cache

import (
	"github.com/onurbaran/easy-cache/pkg/event"
	"github.com/onurbaran/easy-cache/pkg/serializer"
	"time"
)

type EvictionPolicy interface {
	Evict(c *Cache)
}

type Config struct {
	BaseTTL        time.Duration
	MaxItems       int
	EvictionPolicy EvictionPolicy
	TTLOverrides   map[string]time.Duration
	EventManager   *event.EventManager
	Serializer     serializer.Serializer
	NumShards      int
}

func DefaultConfig() *Config {
	return &Config{
		BaseTTL:        10 * time.Minute,
		MaxItems:       100,
		EvictionPolicy: &LRUEviction{},
		TTLOverrides:   make(map[string]time.Duration),
		EventManager:   nil,
		Serializer:     nil,
		NumShards:      8,
	}
}
