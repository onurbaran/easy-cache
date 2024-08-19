package cache

import (
	"context"
	"time"
)

type LRUEviction struct{}

func (e *LRUEviction) Evict(c *Cache) {
	var oldestKey string
	var oldestTime = time.Now()

	for key, item := range c.Items() {
		if item.LastAccessed().Before(oldestTime) {
			oldestKey = key
			oldestTime = item.LastAccessed()
		}
	}

	if oldestKey != "" {
		ctx := context.Background()
		c.Delete(ctx, oldestKey)
	}
}
