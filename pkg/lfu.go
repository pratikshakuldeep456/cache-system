package pkg

import (
	"fmt"
	"math"
)

type LFU struct {
	countMap map[string]int
}

func (l *LFU) evict(c *Cache) {
	fmt.Println("lfu policy")
	if len(l.countMap) == 0 {
		fmt.Println("no key to evict")
	}

	minFreq := math.MaxInt
	var freq string
	for key, count := range l.countMap {
		if minFreq > count {
			freq = key
			minFreq = count
		}
	}
	delete(c.storage, freq)
	delete(l.countMap, freq)
	c.capacity--

	fmt.Println(" lfu evicted key is ", freq)

}

func (l *LFU) AddCache(key string) {
	l.countMap[key]++

}
