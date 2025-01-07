package pkg

import "fmt"

type LRU struct {
	usageOrder []string
}

func (l *LRU) evict(c *Cache) {
	fmt.Println("lru policy")
	if len(l.usageOrder) == 0 {
		fmt.Println("no key to evict")
	}

	firstKey := l.usageOrder[0]
	l.usageOrder = l.usageOrder[1:]
	delete(c.storage, firstKey)
	c.capacity--

}

func (l *LRU) AddCache(key string) {
	//if key present remove and move to last
	for i, k := range l.usageOrder {
		if k == key {
			l.usageOrder = append(l.usageOrder[:i], l.usageOrder[i+1:]...)
			break
		}
	}
	l.usageOrder = append(l.usageOrder, key)

}
