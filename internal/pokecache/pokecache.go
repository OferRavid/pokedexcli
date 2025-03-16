package pokecache

import (
	"sync"
	"time"
)

// Cache is a thread-safe in-memory cache that stores key-value pairs
// with automatic expiration based on a specified interval.
type Cache struct {
	cache map[string]cacheEntry
	mutex *sync.Mutex
}

// cacheEntry represents an individual cache item with a creation timestamp.
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

/*
NewCache creates and returns a new Cache instance with an automatic cleanup process.

The cleanup process removes expired cache entries at the specified interval.

Parameters:
- interval: Duration after which cache entries are reaped.

Returns:
- Cache: A new Cache instance.
*/
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mutex: &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

/*
Add inserts a key-value pair into the cache if the key does not already exist.

Parameters:
- key: A string representing the cache key.
- val: A byte slice containing the data to be stored.
*/
func (c *Cache) Add(key string, val []byte) {
	if _, ok := c.Get(key); ok {
		return
	}

	entry := cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
	c.mutex.Lock()
	c.cache[key] = entry
	c.mutex.Unlock()
}

/*
Get retrieves a value from the cache based on the given key.

Parameters:
- key: A string representing the cache key.

Returns:
- []byte: The cached data.
- bool: True if the key exists, false otherwise.
*/
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, exists := c.cache[key]
	return entry.val, exists
}

/*
reapLoop continuously removes expired cache entries at regular intervals.

Parameters:
- interval: Duration specifying how often expired entries are removed.
*/
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

/*
reap removes expired cache entries that exceed the specified interval.

Parameters:
- now: The current time used to determine expiration.
- interval: The duration threshold beyond which entries are removed.
*/
func (c *Cache) reap(now time.Time, interval time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for key, entry := range c.cache {
		if entry.createdAt.Before(now.Add(-interval)) {
			delete(c.cache, key)
		}
	}
}
