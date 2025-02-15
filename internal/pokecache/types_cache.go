package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntry map[string]CacheEntry
	mu         *sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}
