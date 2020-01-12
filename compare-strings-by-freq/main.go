package main

import (
	"fmt"
	"math"
)

func smallestFreq(s string) int {
	sChar := math.MaxInt64
	cMap := make(map[int]int)

	for _, c := range s {
		if sChar > int(c) {
			sChar = int(c)
		}

		cMap[int(c)]++
	}

	return cMap[sChar]
}

// https://leetcode.com/problems/compare-strings-by-frequency-of-the-smallest-character/
func numSmallerByFrequency(queries []string, words []string) []int {
	qFreq := []int{}
	wFreq := []int{}
	// dpWFreq := []int{}

	for _, q := range queries {
		qFreq = append(qFreq, smallestFreq(q))
	}

	result := make([]int, len(qFreq))
	for _, w := range words {
		wFreq = append(wFreq, smallestFreq(w))
	}

	for i, qf := range qFreq {
		// to make this faster, we could use binary search
		for _, wf := range wFreq {
			if qf < wf {
				result[i]++
			}
		}
	}

	return result
}

func main() {
	fmt.Println(numSmallerByFrequency([]string{"bbb", "cc"},
		[]string{"aa", "a", "aaa", "aaaa"}))
}
