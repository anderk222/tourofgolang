package moretypes

import (
	"fmt"
	"strings"
)

func SlicesOfSlice() { // Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("\x1b[33m[SlicesOfSlice]\x1b[0m  %s\n", strings.Join(board[i], " "))
	}
}
