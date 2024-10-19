package pokecache

import (
    //"fmt"
    "time"
    "sync"
)

type Cache struct {
    cacheEntries map[string]cacheEntry
    mu           *sync.Mutex
}

type cacheEntry struct {
    createdAt time.Time
    val       []byte
}

func NewCache(interval time.Duration) Cache {
    c := Cache{ cacheEntries: make(map[string]cacheEntry), mu: &sync.Mutex{}, }

    go c.reapLoop(interval)

    return c
}

func (c *Cache) Add(key string, val []byte) {
    c.mu.Lock()
    defer c.mu.Unlock()

    c.cacheEntries[key] = cacheEntry{ createdAt: time.Now().UTC(), val: val }
}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()

    val, ok := c.cacheEntries[key]

    return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    
    for range ticker.C {
        c.reap(time.Now().UTC(), interval)
    }
}

func (c *Cache) reap(now time.Time, last time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()

    for k, v := range c.cacheEntries {
        if v.createdAt.Before(now.Add(-last)) {
            delete(c.cacheEntries, k)
        }
    }
}
