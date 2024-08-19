package cache

import (
	"time"
)

type Item struct {
	key          string
	data         interface{}
	accessCount  int
	lastAccessed time.Time
	ttl          time.Duration
	priority     int
}

func NewItem(key string, data interface{}, priority int, ttl time.Duration) *Item {
	return &Item{
		key:          key,
		data:         data,
		accessCount:  0,
		lastAccessed: time.Now(),
		ttl:          ttl,
		priority:     priority,
	}
}

func (item *Item) IncrementAccessCount() *Item {
	item.accessCount++
	return item
}

func (item *Item) RefreshTTL() *Item {
	item.lastAccessed = time.Now()
	return item
}

func (item *Item) IsExpired() bool {
	return time.Since(item.lastAccessed) > item.ttl
}

func (item *Item) GetKey() string {
	return item.key
}

func (item *Item) GetData() interface{} {
	return item.data
}

func (item *Item) AccessCount() int {
	return item.accessCount
}

func (item *Item) LastAccessed() time.Time {
	return item.lastAccessed
}
