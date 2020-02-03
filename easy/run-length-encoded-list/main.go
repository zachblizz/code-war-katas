package main

import "fmt"

func decompressRLElist(nums []int) []int {
	list := []int{}

	for i := 0; i < len(nums); i++ {
		freq := nums[i]
		i++
		val := nums[i]

		for j := 0; j < freq; j++ {
			list = append(list, val)
		}

	}

	return list
}

func main() {
	fmt.Println(decompressRLElist([]int{1, 2, 3, 4}))

}
