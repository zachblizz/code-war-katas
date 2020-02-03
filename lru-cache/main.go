package main

import "fmt"

// lc - https://leetcode.com/problems/lru-cache/

// LRUCache - the cache
type LRUCache struct {
	cap      int
	lruTime  int         // the current time of the LRU
	lruLen   int         // length of the cache...
	lruCache map[int]int // holds the times for the keys...
	cache    map[int]int
}

// Constructor - creates a new LRUCache
func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, -1, 0, make(map[int]int), make(map[int]int)}
}

// Get - gets the key from the cache if preset, -1 if not
// also should update everyone's lruTime
func (l *LRUCache) Get(key int) int {
	if _, ok := l.lruCache[key]; ok {
		newTimes := 1
		l.lruCache[key] = 0

		for k := range l.lruCache {
			if k != key {
				l.lruCache[k] = newTimes
				newTimes++
			}
		}
	}

	if val, ok := l.cache[key]; ok {
		return val
	}

	return -1
}

// Put - puts the value in the cache at key
func (l *LRUCache) Put(key int, value int) {
	l.lruTime++

	for k := range l.lruCache {
		l.lruCache[k]++
	}

	if l.lruLen < l.cap {
		l.lruLen++
	} else {
		var remK int
		maxTime := -1

		for k, ct := range l.lruCache {
			if maxTime < ct {
				maxTime = ct
				remK = k
			}
		}
		fmt.Printf("going to remove %v\n", remK)

		// remove the key from the cache
		delete(l.cache, remK)
		delete(l.lruCache, remK)
	}

	l.lruCache[key] = 0
	l.cache[key] = value
}

func tryTwo() {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}

// need - [null,null,null,null,null,4,3,2,-1,null,-1,2,3,-1,5]
func tryThree() {
	cache := Constructor(3)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)
	fmt.Println(cache.Get(4)) // 4
	fmt.Println(cache.Get(3)) // 3
	fmt.Println(cache.Get(2)) // 2
	fmt.Println(cache.Get(1)) // -1
	fmt.Println("-------")

	cache.Put(5, 5)
	fmt.Println(cache.Get(1)) // -1
	fmt.Println(cache.Get(2)) // 2
	fmt.Println(cache.Get(3)) // 3
	fmt.Println(cache.Get(4)) // -1
	fmt.Println(cache.Get(5)) // 5
}

func tryFour() {
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3)
	cache.Put(4, 1)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
}

func main() {
	// tryTwo()
	fmt.Println("====================")
	tryThree()
	// tryFour()
}
