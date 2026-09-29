package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   sync.Mutex{},
	}
	go c.reapLoop(interval)
	return &c
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mux.Lock()
		for key, entry := range c.cache {
			duration := time.Since(entry.createdAt)
			if duration > interval {
				delete(c.cache, key)
			}
		}
		c.mux.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.cache[key]
	if ok {
		return entry.val, true
	}
	return nil, false
}
