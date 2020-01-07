package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals[:], func(i, j int) bool {
		for x := range intervals[i] {
			if intervals[i][x] == intervals[j][x] {
				continue
			}
			return intervals[i][x] < intervals[j][x]
		}
		return false
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
		[]int{1,4},
		[]int{0,1},
	}

	// intervals = [][]int{
	// 	[]int{1,3},
	// 	[]int{2,8},
	// 	[]int{8,10},
	// 	[]int{10,18},
	// }

	// intervals = [][]int{
	// 	[]int{1,4},
	// 	[]int{0,0},
	// }
	
	intervals = [][]int{
		[]int{2,3},
		[]int{4,5},
		[]int{6,7},
		[]int{8,9},
		[]int{1,10},
	}

	fmt.Println(merge(intervals))
}
