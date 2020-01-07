package main

import (
	"fmt"
	"math"
	// "sort"
)

// kata - https://www.codewars.com/kata/554f76dca89983cc400000bb

// Solequa - gives back a list of (x, y) values
// that suffice the equation x^2-4y^2 = n

// x^2 = n + 4y^2
// x^2 - n = 4y^2
// (x^2 - n)/4 = y^2
func Solequa(n int) [][]int {
	result := [][]int{}

	for x := n; x > 0; x-- {
		xSquared := x * x

		if xSquared-n >= 0 {
			ySquared := (xSquared - n) / 4
			y := math.Sqrt(float64(ySquared))
			
			if xSquared-(4*ySquared) == n && math.Trunc(y) == y {
				result = append(result, []int{x, int(y)})
			}
		}
	}

	return result
}

func main() {
	results := Solequa(900000009)

	fmt.Println(len(results))
	for _, r := range results {
		fmt.Println(r)
	}
}
