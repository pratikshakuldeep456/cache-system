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

	c.capacity--
	l.countMap[freq]--
	fmt.Println(" lfu evicted key is ", freq)
	fmt.Println(" lfu storage is ", l.countMap)

}

func (l *LFU) AddCache(key string) {
	if l.countMap == nil {
		l.countMap = make(map[string]int)
	}
	fmt.Println("djifodf")
	if _, j := l.countMap[key]; j {
		l.countMap[key]++

	} else {
		l.countMap[key] = 1
	}

}
