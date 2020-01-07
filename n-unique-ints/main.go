package main

import "fmt"

// https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero
func main() {
	// every positive number you add, just add the negative number...
	fmt.Println(sumZero(5))
}

func sumZero(n int) []int {
	res := make([]int, n)
	start := 0
	end := n-1

	for ; start < n-1 && start != end; {
		res[end] = n
		res[start] = n*-1

		n--
		start++
		end--
	}
	

	return res
}
