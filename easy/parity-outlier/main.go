package main

// kata - https://www.codewars.com/kata/5526fc09a1bbd946250002dc/train/go

// FindOutlier - finds the outlier within the given array
// Def Outlier - if all ints but one are odd than the one that is not odd is the outlier
// 				 if all ints but one are even than the one that is not even is the outlier
func FindOutlier(integers []int) int {
	oddCount := 0
	evenCount := 0

	for _, i := range integers {
		if i%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	for _, i := range integers {
		isEven := i%2 == 0

		if evenCount == 1 && isEven {
			return i
		} else if oddCount == 1 && !isEven {
			return i
		}
	}

	return 0
}
