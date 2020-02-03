package main

import "fmt"

func sinkTheShip(board *[][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(*board) || j >= len((*board)[0]) || (*board)[i][j] == 46 {
		return
	}

	(*board)[i][j] = 46
	sinkTheShip(board, i-1, j)
	sinkTheShip(board, i+1, j)
	sinkTheShip(board, i, j-1)
	sinkTheShip(board, i, j+1)
}

func countBattleships(board [][]byte) int {
	count := 0

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			if board[row][col] == 88 {
				count++
				sinkTheShip(&board, row, col)
				// printBoard(board)
			}
		}
	}

	return count
}

// func printBoard(board [][]byte) {
// 	for _, rows := range board {
// 		for _, cols := range rows {
// 			fmt.Printf("%s", string(cols))
// 		}
// 		fmt.Println()
// 	}

// 	fmt.Println("\n")
// }

func main() {
	board := [][]byte{
		[]byte{88, 46, 46, 88},
		[]byte{46, 46, 46, 88},
		[]byte{46, 46, 46, 88},
	}

	fmt.Println(countBattleships(board))
}
