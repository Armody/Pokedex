package pokecache

import "time"

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cacheEntry: make(map[string]CacheEntry),
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheEntry[key] = CacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, exists := c.cacheEntry[key]; exists {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	for {
		time.Sleep(interval)

		c.mu.Lock()
		for key, entry := range c.cacheEntry {
			if time.Since(entry.createdAt) >= interval {
				delete(c.cacheEntry, key)
			}
		}
		c.mu.Unlock()
	}
}
