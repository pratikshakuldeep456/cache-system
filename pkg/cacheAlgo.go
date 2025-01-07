package pkg

type Cache struct {
	storage     map[string]string
	algo        EvictionAlgo
	capacity    int
	maxCapacity int
}

type EvictionAlgo interface {
	evict(c *Cache)
	AddCache(key string)
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

func (c *Cache) Get(key string) string {

	val := c.storage[key]
	c.algo.AddCache(key)

	return val

}
func (c *Cache) Put(key string, val string) {
	//delete(c.storage, key)

	if c.capacity == c.maxCapacity {
		c.algo.evict(c)
	}
	c.storage[key] = val
	c.capacity++
	c.algo.AddCache(key)
}

// func (c *Cache) AddCache(value string, key string) {
// 	if c.capacity == c.maxCapacity {
// 		c.algo.evict(c)
// 	}
// 	c.capacity++
// 	c.storage[key] = value

// }

// func (c *Cache) DeleteCache(key string) {
// 	delete(c.storage, key)
// }
