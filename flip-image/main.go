package main

import "fmt"

// https://leetcode.com/problems/flipping-an-image/

func swapIndicies(row []int, i, j int) {
	tmp := row[j]
	row[j] = row[i]
	row[i] = tmp
}

func reverseRow(row []int) []int {	
	end := len(row)-1

	for start := 0; start < end; start++ {
		swapIndicies(row, start, end)
		end--
	}

	return row
}

func invertRow(row []int) []int {
	for i := 0; i < len(row); i++ {
		row[i] ^= 1
	}

	return row
}

func flipAndInvertImage(A [][]int) [][]int {
	for i, row := range A {
		fmt.Println(row)
		r := reverseRow(row)
		fmt.Println("reverse: ", r)
		
		A[i] = invertRow(r)
	}
	return A
}

func main () {
	A := [][]int{
		{1,1,0,0},
		{1,0,0,1},
		{0,1,1,1},
		{1,0,1,0},
	}
	flipAndInvertImage(A)
	fmt.Println(A)
}

