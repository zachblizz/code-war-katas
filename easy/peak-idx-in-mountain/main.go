package main

// leet - https://leetcode.com/problems/peak-index-in-a-mountain-array

func peakIndexInMountainArray(A []int) int {
	isPeak := false
	p := A[0]

	for i := 1; i < len(A); i++ {
		c := A[i]

		if p < c && !isPeak {
			isPeak = true
		} else if p > c && isPeak {
			return i-1
		}

		p = c
	}

	return 0
}

func main() {

}
