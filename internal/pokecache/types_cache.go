package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
