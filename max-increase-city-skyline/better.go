package main

func maxIncreaseKeepingSkylineBetter(grid [][]int) int {
    maxRow := make([]int, len(grid))
    maxColumn := make([]int, len(grid[0]))

    for i := range maxRow {
        row := grid[i]
        maxRow[i] = max(row...)
    }

    for i := range maxColumn {
        columnVals := make([]int, len(maxColumn))
        for j := range grid {
            columnVals = append(columnVals, grid[j][i])
        }
        maxColumn[i] = max(columnVals...)
    }

    sum := 0
    for i := range grid {
        for j, item := range grid[i] {
            skylineMax := min(maxRow[i], maxColumn[j])
            sum += skylineMax - item
        }
    }

    return sum
}

func max(values ...int) int {
    currMax := 0
    for _, value := range values {
        if value > currMax {
            currMax = value
        }
    }

    return currMax
}

func min(values ...int) int {
    currMin := 100
    for _, value := range values {
        if value < currMin {
            currMin = value
        }
    }

    return currMin
}
