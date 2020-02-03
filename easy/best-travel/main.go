package main

// kata - https://www.codewars.com/kata/55e7280b40e1c4a06d0000aa/train/go

// ChooseBestSum_GARBAGE - finds the largest number of k values in ls that equal closest to t
func ChooseBestSum_GARBAGE(t, k int, ls []int) int {
	if len(ls) < k {
		return -1
	}

	largest := -1

	for i := 0; i < len(ls); i++ {
		towns := ls[i]
		seen := 0

		for j := 0; j < len(ls); j++ {
			if ls[j] != ls[i] && seen <= k && towns <= t {
				seen++
				towns += ls[j]
			}

			if seen == k && towns <= t {
				largest = towns
			} else if towns > t {
				seen = 0
				towns = ls[i]
			}
		}
	}

	return largest
}

// ChooseBestSum - finds the largest number of k values in ls that equal closest to t
func ChooseBestSum(t, k int, ls []int) int {
	outerbest := -1

	for i, d := range ls {
		// not enough remaining values for this d to work
		if len(ls) < k {
			continue
		}

		// recursively choose best from t-d, until final level k=1
		if k > 1 {
			innerbest := ChooseBestSum(t-d, k-1, ls[i+1:])

			// if no best available at lower level, this d cant work
			if innerbest < 0 {
				continue
			}

			d += innerbest
		}

		if d <= t && d > outerbest {
			outerbest = d
		}
	}

	return outerbest
}
