package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	entries map[string]cacheEntry
	mutex sync.Mutex
}

type cacheEntry struct{
	createdAt time.Time
	val []byte
}