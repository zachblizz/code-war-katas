package main

import "fmt"

// kata - https://www.codewars.com/kata/5547cc7dcad755e480000004/train/go

// RemovNb - removes the (a,b) pair from the list up to m
// that's product equals the sum of the rest of the list
func RemovNb(m uint64) [][2]uint64 {
	total := summit1G(m)

	for i := uint64(1); i <= m; i++ {
		for j := uint64(1); j <= m; j++ {
			product := i*j
			newSum := total-(i+j)

			if product == newSum {
				return [][2]uint64{{i,j},{j,i}}
			}
		}
	}

	return nil
}

func summit1G(m uint64) uint64 {
	sum := uint64(0)

	for i := uint64(1); i <= m; i++ {
		sum += i
	}

	return sum
}

func main() {
	fmt.Println(RemovNb(1000003))
}
