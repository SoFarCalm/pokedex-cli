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
	entries map[string]CacheEntry
	mu      sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	ticker := time.NewTicker(interval)

	newCache := &Cache{
		entries: make(map[string]CacheEntry),
	}

	go func() {
		for range ticker.C {
			newCache.Cleanup(interval)
		}
	}()

	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if entry, exists := c.entries[key]; exists {
		return entry.val, true
	}

	return nil, false
}

func (c *Cache) Cleanup(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.entries {
		if time.Since(entry.createdAt) > interval {
			delete(c.entries, key)
		}
	}
}
