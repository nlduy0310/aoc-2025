package cache

import "sync"

type Cache[V any] struct {
	mutex sync.RWMutex
	map_  map[uint64]V
}

func New[V any]() Cache[V] {
	return Cache[V]{
		map_: make(map[uint64]V),
	}
}

func (c *Cache[V]) Get(key uint64) (V, bool) {
	c.mutex.RLock()
	v, ok := c.map_[key]
	c.mutex.RUnlock()
	return v, ok
}

func (c *Cache[V]) Set(key uint64, val V) {
	c.mutex.Lock()
	c.map_[key] = val
	c.mutex.Unlock()
}
