package main

import "sort"

// func heightChecker(heights []int) int {
// 	res := 0
// 	prev := 0
// 	next := 0

// 	for i, curr := range heights {
// 		if (i+1 < len(heights)) {
// 			next = heights[i+1]
// 		}

// 		cGtn := i > 0 && curr > next

// 		if cGtn {
// 			res++
// 		}

// 		if i > 0 && curr < prev && next >= prev {
// 			res++
// 		}

// 		prev = curr
// 	}

// 	return res
// }

func heightChecker(heights []int) int {
	res := 0
	c := make([]int, len(heights))
	copy(c, heights)

	sort.Ints(c)
	for i, h := range heights {
		if h != c[i] {
			res++
		}
	}

	return res
}
