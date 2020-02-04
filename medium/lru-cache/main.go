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
		l.lruCache[key] = 0

		for k := range l.lruCache {
			if k != key {
				l.lruCache[k]++
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

	if _, ok := l.lruCache[key]; !ok && l.lruLen < l.cap {
		l.lruLen++
	} else if _, ok := l.lruCache[key]; !ok {
		var remK int
		maxTime := -1

		for k, ct := range l.lruCache {
			if maxTime < ct {
				maxTime = ct
				remK = k
			}
		}
		// fmt.Printf("going to remove %v\n", remK)

		// remove the key from the cache
		delete(l.cache, remK)
		delete(l.lruCache, remK)
	}

	l.lruCache[key] = 0
	l.cache[key] = value
}

func setup(ops []string, vals [][]int) {
	var cache LRUCache
	out := []int{}

	for i, op := range ops {
		switch op {
		case "LRUCache":
			cache = Constructor(vals[i][0])
			break
		case "put":
			cache.Put(vals[i][0], vals[i][1])
			break
		case "get":
			out = append(out, cache.Get(vals[i][0]))
			break
		}
	}

	fmt.Println(out)
}

func main() {
	// -1,-1,2,6
	ops := []string{"LRUCache", "get", "put", "get", "put", "put", "get", "get"}
	vals := [][]int{[]int{2}, []int{2}, []int{2, 6}, []int{1}, []int{1, 5}, []int{1, 2}, []int{1}, []int{2}}

	// // -1,3
	// ops = []string{"LRUCache", "put", "put", "put", "put", "get", "get"}
	// vals = [][]int{[]int{2}, []int{2, 1}, []int{1, 1}, []int{2, 3}, []int{4, 1}, []int{1}, []int{2}}

	// // 1,-1,-1,3,4
	//	ops = []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"}
	//	vals = [][]int{[]int{2}, []int{1, 1}, []int{2, 2}, []int{1}, []int{3, 3}, []int{2}, []int{4, 4}, []int{1}, []int{3}, []int{4}}

	setup(ops, vals)
}
