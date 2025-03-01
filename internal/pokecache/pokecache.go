package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache    map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

func NewCache(someInterval time.Duration) *Cache {
	c := Cache{}
	c.cache = make(map[string]cacheEntry)
	c.interval = someInterval
	go c.reapLoop()
	return &c
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	newCacheEntry := cacheEntry{}
	newCacheEntry.createdAt = time.Now()
	newCacheEntry.val = value
	c.cache[key] = newCacheEntry
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	item, exists := c.cache[key]
	c.mu.Unlock()
	if exists {
		return item.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	timeTicker := time.NewTicker(c.interval)
	defer timeTicker.Stop()

	for range timeTicker.C {
		c.mu.Lock()
		for key, entry := range c.cache {
			if entry.createdAt.Add(c.interval).Before(time.Now()) {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}
