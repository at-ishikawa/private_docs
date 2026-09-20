package main

func getCellCandidates(board [][]byte, y int, x int) []byte {
	cellCandidates := make(map[byte]bool)
	for i := 0; i < len(board); i++ {
		// check values in the same column
		cellCandidates[board[i][x]] = true
		// check values in the same row
		cellCandidates[board[y][i]] = true
	}

	// check values in a sub-boxes
	subBoxY := (y / 3)
	subBoxX := (x / 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			cell := board[subBoxY*3+i][subBoxX*3+j]
			cellCandidates[cell] = true
		}
	}

	result := make([]byte, 0)
	for i := '1'; i <= '9'; i++ {
		if _, ok := cellCandidates[byte(i)]; ok {
			continue
		}
		result = append(result, byte(i))
	}
	return result
}

func trySudoku(board [][]byte, startY, startX int) bool {
	for y := startY; y < len(board); y++ {
		x := 0
		if y == startY {
			x = startX
		}
		for ; x < len(board[y]); x++ {
			cell := board[y][x]
			if cell != '.' {
				continue
			}

			candidates := getCellCandidates(board, y, x)
			for i := 0; i < len(candidates); i++ {
				board[y][x] = candidates[i]
				// backtrack
				if trySudoku(board, y, x+1) {
					return true
				}

				board[y][x] = '.'
			}
			// no candidate was found
			return false
		}
	}

	return true
}

func solveSudoku(board [][]byte) {
	trySudoku(board, 0, 0)
}

/*
type Sudoku struct {
	board           [][]byte
	boardCandidates [][][]byte
}

func getCellCandidates(board [][]byte, y int, x int) []byte {
	cellCandidates := make(map[byte]bool)
	// check values in the same column
	for i := 0; i < len(board); i++ {
		cell := board[i][x]
		cellCandidates[cell] = true
	}

	// check values in the same row
	for i := 0; i < len(board[0]); i++ {
		cell := board[y][i]
		cellCandidates[cell] = true
	}

	// check values in a sub-boxes
	subBoxY := (y / 3)
	subBoxX := (x / 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			cell := board[subBoxY*3+i][subBoxX*3+j]
			cellCandidates[cell] = true
		}
	}

	result := make([]byte, 0)
	for i := '1'; i <= '9'; i++ {
		if _, ok := cellCandidates[byte(i)]; ok {
			continue
		}
		result = append(result, byte(i))
	}
	return result
}

func newBoardCandidates(board [][]byte) [][][]byte {
	candidates := make([][][]byte, len(board))
	for i := 0; i < len(board); i++ {
		candidates[i] = make([][]byte, len(board[i]))
	}

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if board[y][x] != '.' {
				continue
			}

			candidates[y][x] = getCellCandidates(board, y, x)
		}
	}
	return candidates
}

func (sudoku Sudoku) findValue(y, x int) byte {
	if len(sudoku.boardCandidates[y][x]) == 1 {
		return sudoku.boardCandidates[y][x][0]
	}

	cellCandidates := make([]byte, len(sudoku.boardCandidates[y][x]))
	copy(cellCandidates, sudoku.boardCandidates[y][x])

	checkCell := func(y2, x2 int) {
		if y == y2 && x == x2 {
			return
		}

		otherCellCandidates := sudoku.boardCandidates[y2][x2]
		for i := range otherCellCandidates {
			otherCellCandidate := otherCellCandidates[i]
			index := slices.Index(cellCandidates, otherCellCandidate)
			if index == -1 {
				continue
			}
			cellCandidates = slices.Delete(cellCandidates, index, index+1)
		}
	}
	for i := 0; i < len(sudoku.board); i++ {
		checkCell(i, x)
	}
	for i := 0; i < len(sudoku.board[y]); i++ {
		checkCell(y, i)
	}
	subBoxY := (y / 3)
	subBoxX := (x / 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			checkCell(subBoxY*3+i, subBoxX*3+j)
		}
	}

	var result byte
	if len(cellCandidates) == 1 {
		result = cellCandidates[0]
	}

	// fmt.Printf("%v: (%d, %d): %+v %+v\n", result, y, x, sudoku.boardCandidates[y][x], cellCandidates)
	return result
}

func solveSudoku(board [][]byte) {
	sudoku := Sudoku{
		board: board,
	}

	isProgress := true
	for isProgress {
		sudoku.boardCandidates = newBoardCandidates(sudoku.board)
		isProgress = false
		for y := 0; y < len(board); y++ {
			for x := 0; x < len(board[y]); x++ {
				cell := board[y][x]
				if cell != '.' {
					continue
				}

				value := sudoku.findValue(y, x)
				if value != 0 {
					sudoku.board[y][x] = value
				} else {
					isProgress = true
				}
			}
		}

		fmt.Printf("---------------Board--------------\n")
		for y := 0; y < len(sudoku.board); y++ {
			fmt.Printf("%+v\n", string(sudoku.board[y]))
		}
	}
}
*/
