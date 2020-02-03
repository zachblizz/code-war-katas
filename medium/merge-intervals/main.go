package main

// leet code - https://leetcode.com/problems/merge-intervals/

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals[:], func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	return helper(intervals, [][]int{}, intervals[0], intervals[0], 0)
}

func helper(intervals, ret [][]int, interval, lastMerged []int, idx int) [][]int {
	if lastMerged[1] >= interval[0] && lastMerged[0] <= interval[1] {
		tmp := interval[1]

		if lastMerged[0] > interval[0] {
			lastMerged[0] = interval[0]
		}

		if lastMerged[1] < interval[1] {
			lastMerged[1] = tmp
		}
	} else {
		ret = append(ret, lastMerged)
		lastMerged = interval
	}

	idx++
	if idx < len(intervals) {
		return helper(intervals, ret, intervals[idx], lastMerged, idx)
	}

	ret = append(ret, lastMerged)
	return ret
}

func main() {
	intervals := [][]int{
		[]int{1, 4},
		[]int{0, 1},
	}

	fmt.Println(merge(intervals))
}
