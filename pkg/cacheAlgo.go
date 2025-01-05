package pkg

import "fmt"

type LRU struct {
}

func (l *LRU) evict(c *Cache) {
	fmt.Println("lru policy")
}

type LFU struct {
}

func (l *LFU) evict(c *Cache) {
	fmt.Println("lfu policy")
}

type Cache struct {
	storage     map[string]string
	algo        EvictionAlgo
	capacity    int
	maxCapacity int
}
type EvictionAlgo interface {
	evict(c *Cache)
}

func InitCache(e EvictionAlgo) *Cache {
	return &Cache{
		storage:     make(map[string]string),
		algo:        e,
		capacity:    0,
		maxCapacity: 10,
	}
}

func (c *Cache) SetCacheAlgo(e EvictionAlgo) {
	c.algo = e
}

func (c *Cache) AddCache(value string, key string) {
	if c.capacity == c.maxCapacity {
		c.algo.evict(c)
	}
	c.capacity++
	c.storage[key] = value

}

func (c *Cache) DeleteCache(key string) {
	delete(c.storage, key)
}
