package internal

// var queens [][]int

// Time complexity: O(N^3)
//
//	N*N: To search board
//	(N*N)*N*8: To mark a board
//
// Space complexity: O(N^2)
//
//	N*N: board
func totalNQueens(n int) int {
	board := make([][]bool, n)
	for i := 0; i < n; i++ {
		board[i] = make([]bool, n)
	}

	// queens = make([][]int, 0, len(board))
	result := 0
	for i := 0; i < n; i++ {
		result += searchNBoard(board, 0, i, 0, n)
	}
	return result
}

func updateBoard(board [][]bool, row int, column int, isQueenPut bool) [][2]int {
	n := len(board)

	directions := [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},

		// right down
		{1, -1},
		{-1, -1},

		// right up
		{-1, 1},
		{1, 1},
	}

	board[row][column] = isQueenPut
	result := [][2]int{
		{row, column},
	}
	for _, direction := range directions {
		for i := 1; i < n; i++ {
			nextRow := i*direction[0] + row
			nextColumn := i*direction[1] + column
			if nextRow < 0 || n <= nextRow || nextColumn < 0 || n <= nextColumn {
				break
			}
			if board[nextRow][nextColumn] == isQueenPut {
				continue
			}
			board[nextRow][nextColumn] = isQueenPut
			result = append(result, [2]int{nextRow, nextColumn})
		}
	}
	return result
}

func searchNBoard(board [][]bool, row int, column int, queenCount int, n int) int {
	if row < 0 || len(board) <= row || column < 0 || len(board[0]) <= column {
		return 0
	}

	if board[row][column] {
		return 0
	}
	if queenCount+1 == n {
		// for _, queen := range queens {
		// 	fmt.Printf("(%d, %d), ", queen[0], queen[1])
		// }
		// fmt.Printf("(%d, %d)\n", row, column)
		return 1
	}

	marked := updateBoard(board, row, column, true)
	// queens = append(queens, []int{row, column})
	result := 0
	for i := 0; i < n; i++ {
		result += searchNBoard(board, row+1, i, queenCount+1, n)
	}
	for _, cell := range marked {
		r := cell[0]
		c := cell[1]
		board[r][c] = false
	}
	// queens = queens[:len(queens)-1]
	return result
}
