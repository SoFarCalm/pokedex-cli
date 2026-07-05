package pokecache

import (
	"sync"
	"time"
)

// The NewCache function creates a new cache instance and starts an automatic cleanup process.

// Here is how it works step by step:

// It creates a ticker
// The line:
// ticker := time.NewTicker(interval)
// creates a timer-like object that fires at regular intervals based on the interval argument.
// It initializes the cache struct
// The code builds a new Cache value:
// entries: make(map[string]CacheEntry)
// This creates an empty map that will store cached data keyed by string.
// It starts a background goroutine
// The function launches:
// go func() {
// for range ticker.C {
// newCache.Cleanup(interval)
// }
// }()
// This means a separate lightweight thread (goroutine) runs forever in the background.
// On every tick, it calls Cleanup(interval), which removes old entries from the cache.
// It returns the cache
// The constructor returns the newly created cache pointer:
// return newCache
// So in practical terms, NewCache does two things:

// sets up the internal storage for cached entries
// starts a background process that periodically removes expired entries
// Why this matters:

// The cache does not need manual cleanup from the caller.
// Old data is automatically removed after the specified interval, helping prevent stale data from lingering indefinitely.
// One subtle detail:

// The cleanup loop runs continuously for the lifetime of the cache.
// It uses the same interval passed into NewCache, so the expiration policy is consistent.
// In short, NewCache is essentially “initialize the cache and start its automatic expiration mechanism.”

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
