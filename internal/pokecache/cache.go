package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache() (*Cache, error) {
	interval := (5 * time.Second)

	cache := &Cache{}

	return cache, nil
}

func (c *Cache) Add(key string, val []byte) error {

	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	var val []byte

	return val, true
}

func (c *Cache) readLoop() {

}
