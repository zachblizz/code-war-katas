package main

import (
	"fmt"
	"sort"
)

func smallestFreq(s string) int {
	sChar := 'z'
	cMap := make(map[rune]int)

	for _, c := range s {
		if sChar > c {
			sChar = c
		}

		cMap[c]++
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
	sort.Ints(wFreq)

	for i, qf := range qFreq {
		// to make this faster, we could use binary search
		result[i] = binarySearch(qf, wFreq)
		// for _, wf := range wFreq {
		// 	if qf < wf {
		// 		result[i]++
		// 	}
		// }
	}

	return result
}

func binarySearch(num int, wFreq []int) int {
	start := 0
	end := len(wFreq) - 1

	for start <= end {
		mid := (start + end) / 2

		if wFreq[mid] <= num {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return len(wFreq) - start
}

func main() {
	fmt.Println(numSmallerByFrequency([]string{"bbb", "cc"},
		[]string{"aa", "a", "aaa", "aaaa"}))
}
