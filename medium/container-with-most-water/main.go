package main

// lc - https://leetcode.com/problems/container-with-most-water/

func maxArea(height []int) int {
	res := 0
	s := 0
	e := len(height) - 1

	for s < e {
		area := min(&height[s], &height[e]) * (e - s)
		res = max(&res, &area)

		if height[s] < height[e] {
			s++
		} else {
			e--
		}
	}

	return res
}

func min(a, b *int) int {
	if *a < *b {
		return *a
	}

	return *b
}

func max(a, b *int) int {
	if *a > *b {
		return *a
	}

	return *b
}

func main() {

}
