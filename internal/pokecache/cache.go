package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]CacheEntry
	mu    *sync.Mutex
}

// New Cache
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]CacheEntry),
		mu:    &sync.Mutex{},
	}

	go c.CLeanupLoop(interval)

	return c
}

// Add
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = CacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

// Get
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if entry, exists := c.cache[key]; exists {
		return entry.val, true
	}

	return nil, false
}

func (c *Cache) CLeanupLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.Cleanup(interval)
	}
}

func (c *Cache) Cleanup(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.cache {
		if time.Since(entry.createdAt) > interval {
			delete(c.cache, key)
		}
	}
}
