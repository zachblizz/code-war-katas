package main

import "fmt"

func main() {
	a := [][]int{
		{0, 1},
		{1, 1},
	}

	// fmt.Println(oddCells(2, 3, a))
	fmt.Println(oddCells2(2, 3, a))
}

/*
given 'n' and 'm' which are the dimensions of a matrix
initialized by zeros and given an array indices where
indicies[i] = [ri, ci]. For each pair of [ri, ci] you
have to increment all cells in row ri and column ci by 1.

Return the *number of cells* with odd values in the
matrix after applying the increment to all indices.
*/
func oddCells(n int, m int, indecies [][]int) int {
	rows := make([]int, n)
	cols := make([]int, m)
	oddr := 0
	oddc := 0

	for _, updates := range indecies {
		rows[updates[0]]++
		cols[updates[1]]++
	}

	for _, r := range rows {
		if r%2 == 1 {
			oddr++
		}
	}

	for _, c := range cols {
		if c%2 == 1 {
			oddc++
		}
	}

	return oddr*m + oddc*n - 2*oddr*oddc
}

func oddCells2(n int, m int, indices [][]int) int {
	t := make([][]int, n)
	ret := 0

	for i := range t {
		t[i] = make([]int, m)
	}

	for _, updates := range indices {
		helper(updates, t)
	}

	for _, x := range t {
		for _, v := range x {
			if v%2 == 1 {
				ret++
			}
		}
	}

	return ret
}

func helper(updates []int, newA [][]int) {
	ri := updates[0]
	ci := updates[1]

	for i := range newA[ri] {
		newA[ri][i]++
	}

	for i := range newA {
		newA[i][ci]++
	}
}
