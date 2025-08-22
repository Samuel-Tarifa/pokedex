package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache{
	c:=Cache{
		entries: map[string]cacheEntry{},
		mutex:sync.Mutex{},
	}
	go c.reapLoop(interval)
	return &c
}

func (c *Cache) Add(key string,val []byte){
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.entries[key]=cacheEntry{
		val:val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte,bool){
	c.mutex.Lock()
	defer c.mutex.Unlock()
	val,ok:=c.entries[key]
	if !ok{
		return nil,false
	}
	return val.val,true
}

func (c *Cache) reapLoop(interval time.Duration){
	ticker:=time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C{

		c.mutex.Lock()
		for k:=range c.entries{
			entry:=c.entries[k]
			if time.Since(entry.createdAt)>interval{
				delete(c.entries,k)
			}
		}

		c.mutex.Unlock()
	}
}