package cache

import (
	"context"
)

type LFUEviction struct{}

func (e *LFUEviction) Evict(c *Cache) {
	var leastKey string
	var leastAccess = int(^uint(0) >> 1)

	for key, item := range c.Items() {
		if item.AccessCount() < leastAccess {
			leastKey = key
			leastAccess = item.AccessCount()
		}
	}

	if leastKey != "" {
		ctx := context.Background()
		c.Delete(ctx, leastKey)
	}
}
