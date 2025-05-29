package main

import (
	"fmt"
)

const size = 3

type Board [size][size]string

func NewBoard() Board {
	var b Board
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			b[i][j] = " "
		}
	}
	return b
}

func (b Board) Print() {
	fmt.Println("  0 1 2")
	for i, row := range b {
		fmt.Print(i, " ")
		for j, cell := range row {
			fmt.Print(cell)
			if j < size-1 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < size-1 {
			fmt.Println("  -+-+-")
		}
	}
}

func (b Board) CheckWinner() (bool, string) {
	for i := 0; i < size; i++ {
		if b[i][0] != " " && b[i][0] == b[i][1] && b[i][1] == b[i][2] {
			return true, b[i][0]
		}
		if b[0][i] != " " && b[0][i] == b[1][i] && b[1][i] == b[2][i] {
			return true, b[0][i]
		}
	}
	if b[0][0] != " " && b[0][0] == b[1][1] && b[1][1] == b[2][2] {
		return true, b[0][0]
	}
	if b[0][2] != " " && b[0][2] == b[1][1] && b[1][1] == b[2][0] {
		return true, b[0][2]
	}
	return false, ""
}

func (b Board) IsFull() bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if b[i][j] == " " {
				return false
			}
		}
	}
	return true
}

func main() {
	board := NewBoard()
	players := []string{"O", "X"}
	turn := 0

	for {
		board.Print()
		fmt.Printf("Player %s, input (row col): ", players[turn])
		var row, col int
		fmt.Scan(&row, &col)
		if row < 0 || row >= size || col < 0 || col >= size || board[row][col] != " " {
			fmt.Println("Invalid move, try again.")
			continue
		}
		board[row][col] = players[turn]

		if win, winner := board.CheckWinner(); win {
			board.Print()
			fmt.Printf("Player %s wins!\n", winner)
			break
		}
		if board.IsFull() {
			board.Print()
			fmt.Println("Draw!")
			break
		}
		turn = 1 - turn
	}
}
