package main

// problem - https://leetcode.com/problems/max-increase-to-keep-city-skyline

// Max - holds the row max and the col max
type Max struct {
	rowMax int
	colMax int
}

// MaxIncreaseKeepingSkyline - returns the max total sum that
// the height of the buildings can be increased
func MaxIncreaseKeepingSkyline(grid [][]int) int {
	origGridSum := calcGrid(grid)
	grid = reconstructGrid(grid)
	newGridSum := calcGrid(grid)

	return newGridSum - origGridSum
}

func calcGrid(grid [][]int) int {
	sum := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid); col++ {
			sum += grid[row][col]
		}
	}

	return sum
}

func reconstructGrid(grid [][]int) [][]int {
	buildings := len(grid)
	maxes := getMaxes(grid)

	for row := 0; row < buildings; row++ {
		for col := 0; col < buildings; col++ {
			b := grid[row][col]
			maxPossible := findMaxPossible(maxes, row, col)

			if b < maxPossible {
				grid[row][col] = maxPossible
			}
		}
	}

	return grid
}

func findMaxPossible(maxes []Max, row, col int) int {
	rowMax := maxes[row].rowMax
	colMax := maxes[col].colMax
	result := rowMax

	if rowMax > colMax {
		result = colMax
	}

	return result
}

func getMaxes(grid [][]int) []Max {
	maxes := make([]Max, len(grid))

	for row := 0; row < len(grid); row++ {
		rowMax := grid[row][0]
		maxes[row] = Max{rowMax, 0}

		for col := 0; col < len(grid); col++ {
			rb := grid[row][col]

			if rowMax < rb {
				rowMax = rb
			}
		}

		maxes[row].rowMax = rowMax
	}

	for row := 0; row < len(grid); row++ {
		colMax := grid[0][row]
		maxes[row].colMax = colMax

		for col := 0; col < len(grid); col++ {
			cb := grid[col][row]

			if colMax < cb {
				colMax = cb
			}
		}

		maxes[row].colMax = colMax
	}

	return maxes
}
