package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]cacheEntry
	interval time.Duration
	mu       sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) (*Cache, error) {

	cache := &Cache{
		cache:    make(map[string]cacheEntry),
		interval: interval,
	}
	go cache.reapLoop()

	return cache, nil
}

func (c *Cache) Add(key string, val []byte) error {
	ce := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = ce
	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, success := c.cache[key]
	if success == false {
		return nil, false
	}

	return val.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	for {
		<-ticker.C
		c.mu.Lock()
		for key, entry := range c.cache {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}

}
