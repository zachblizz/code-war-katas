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

	if l.lruLen < l.cap {
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
	// expected:  -1,19,17,-1,-1,-1,5,-1,12,3,5,5,1,-1,30,5,30,-1,-1,24,18,-1,18,-1,18,-1,4,29,30,12,-1,29,17,22,18,-1,20,-1,18,18,20
	// need:      -1 19 17 -1 -1 -1 5 -1 12 3 5 5 1 -1 -1 5 30 -1 -1 24 18 -1 18 -1 -1 -1 4 29 -1 12 -1 29 17 -1 -1 -1 -1 -1 -1 -1 -1
	ops := []string{"LRUCache", "put", "put", "put", "put", "put", "get", "put", "get", "get", "put", "get", "put", "put", "put", "get", "put", "get", "get", "get", "get", "put", "put", "get", "get", "get", "put", "put", "get", "put", "get", "put", "get", "get", "get", "put", "put", "put", "get", "put", "get", "get", "put", "put", "get", "put", "put", "put", "put", "get", "put", "put", "get", "put", "put", "get", "put", "put", "put", "put", "put", "get", "put", "put", "get", "put", "get", "get", "get", "put", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "get", "get", "get", "put", "put", "put", "get", "put", "put", "put", "get", "put", "put", "put", "get", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "put", "put", "put"}
	vals := [][]int{[]int{10}, []int{10, 13}, []int{3, 17}, []int{6, 11}, []int{10, 5}, []int{9, 10}, []int{13}, []int{2, 19}, []int{2}, []int{3}, []int{5, 25}, []int{8}, []int{9, 22}, []int{5, 5}, []int{1, 30}, []int{11}, []int{9, 12}, []int{7}, []int{5}, []int{8}, []int{9}, []int{4, 30}, []int{9, 3}, []int{9}, []int{10}, []int{10}, []int{6, 14}, []int{3, 1}, []int{3}, []int{10, 11}, []int{8}, []int{2, 14}, []int{1}, []int{5}, []int{4}, []int{11, 4}, []int{12, 24}, []int{5, 18}, []int{13}, []int{7, 23}, []int{8}, []int{12}, []int{3, 27}, []int{2, 12}, []int{5}, []int{2, 9}, []int{13, 4}, []int{8, 18}, []int{1, 7}, []int{6}, []int{9, 29}, []int{8, 21}, []int{5}, []int{6, 30}, []int{1, 12}, []int{10}, []int{4, 15}, []int{7, 22}, []int{11, 26}, []int{8, 17}, []int{9, 29}, []int{5}, []int{3, 4}, []int{11, 30}, []int{12}, []int{4, 29}, []int{3}, []int{9}, []int{6}, []int{3, 4}, []int{1}, []int{10}, []int{3, 29}, []int{10, 28}, []int{1, 20}, []int{11, 13}, []int{3}, []int{3, 12}, []int{3, 8}, []int{10, 9}, []int{3, 26}, []int{8}, []int{7}, []int{5}, []int{13, 17}, []int{2, 27}, []int{11, 15}, []int{12}, []int{9, 19}, []int{2, 15}, []int{3, 16}, []int{1}, []int{12, 17}, []int{9, 1}, []int{6, 19}, []int{4}, []int{5}, []int{5}, []int{8, 1}, []int{11, 7}, []int{5, 2}, []int{9, 28}, []int{1}, []int{2, 2}, []int{7, 4}, []int{4, 22}, []int{7, 24}, []int{9, 26}, []int{13, 28}, []int{11, 26}}

	// -1,-1,2,6
	// ops = []string{"LRUCache", "get", "put", "get", "put", "put", "get", "get"}
	// vals = [][]int{[]int{2}, []int{2}, []int{2, 6}, []int{1}, []int{1, 5}, []int{1, 2}, []int{1}, []int{2}}

	// // -1,3
	// ops = []string{"LRUCache", "put", "put", "put", "put", "get", "get"}
	// vals = [][]int{[]int{2}, []int{2, 1}, []int{1, 1}, []int{2, 3}, []int{4, 1}, []int{1}, []int{2}}

	// // 1,-1,-1,3,4
	ops = []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"}
	vals = [][]int{[]int{2}, []int{1, 1}, []int{2, 2}, []int{1}, []int{3, 3}, []int{2}, []int{4, 4}, []int{1}, []int{3}, []int{4}}

	setup(ops, vals)
}
