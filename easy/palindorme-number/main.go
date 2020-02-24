package main

import "fmt"

func isPalindrome(x int) bool {
	arr := []int{}
	a := x

	for a > 0 {
		t := a % 10
		arr = append(arr, t)
		a /= 10
	}

	s := 0
	e := len(arr) - 1

	for ; s < e; s++ {
		if arr[s] != arr[e] || arr[s] < 0 || arr[e] < 0 {
			return false
		}
		e--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(-121))
}
