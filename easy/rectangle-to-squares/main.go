package main

import "fmt"

// link - https://www.codewars.com/kata/rectangle-into-squares/train/go

// SquaresInRect - calculates the number of rects in a
// true rectangle ("true" rectangle meaning that the two dimensions are different)
func SquaresInRect(lng, wdth int) []int {
	if lng == wdth {
		return nil
	}

	rect := makeRect(lng, wdth)
	lSquare := getLargestPossibleSquare(lng, wdth)

	fillOutSquare(lSquare, rect)

	return []int{}
}

func makeRect(l, w int) [][]int {
	rect := make([][]int, l)

	for i := range rect {
		rect[i] = make([]int, w)
	}

	return rect
}

func getLargestPossibleSquare(l, w int) int {
	lSquare := l

	if lSquare > w {
		lSquare = w
	}

	return lSquare
}

func fillOutSquare(square int, rect [][]int) {
	sR, sC, eR, eC := getStartRowAndCol(square, rect)

	for row := sR; row < len(rect[0]) && row < eR; row++ {
		for col := sC; col < len(rect[0]) && col < eC; col++ {
			rect[row][col] = 1
		}
	}
}

func getStartRowAndCol(square int, rect [][]int) (int, int, int, int) {
	sR := 0
	sC := 0

	if rect[0][0] == 0 {
		return sR, sC, square, square
	}

	for row := range rect {
		for col := range rect[0] {
			if rect[row][col] != 1 {
				sC = col
				sR = row

				break
			}
		}

		if sR != 0 || sC != 0 {
			break
		}
	}

	eR := sR + square
	eC := sC + square

	return sR, sC, eR, eC
}

func main() {
	rect := [][]int{
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,0},
	}

	fillOutSquare(3, rect)
	printRect(rect)
	fillOutSquare(2, rect)
	printRect(rect)
	fillOutSquare(1, rect)
	printRect(rect)
	fillOutSquare(1, rect)
	printRect(rect)
}

func printRect(rect [][]int) {
	for _, c := range rect {
		fmt.Println(c)
	}
	fmt.Println()
}
