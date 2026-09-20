package main

// Once we found a X< update a neighbor to . if there is X and count

func countBattleships(board [][]byte) int {
	m := len(board)
	n := len(board[0])

	count := 0
	for row := 0; row < m; row++ {
		for column := 0; column < n; column++ {
			if board[row][column] == '.' {
				continue
			}
			if row > 0 && board[row-1][column] == 'X' {
				continue
			}
			if column > 0 && board[row][column-1] == 'X' {
				continue
			}
			count++
		}
	}
	return count
}

func countBattleships2(board [][]byte) int {
	m := len(board)
	n := len(board[0])

	battleshipped := make([][]bool, m)
	for i := 0; i < m; i++ {
		battleshipped[i] = make([]bool, n)
	}

	count := 0
	for row := 0; row < m; row++ {
		for column := 0; column < n; column++ {
			if board[row][column] == '.' {
				continue
			}
			if battleshipped[row][column] {
				continue
			}

			count++
			for i := row; i < m; i++ {
				if board[i][column] == '.' {
					break
				}
				battleshipped[i][column] = true
			}
			// assumption: there is no invalid battleship
			for i := column; i < n; i++ {
				if board[row][i] == '.' {
					break
				}
				battleshipped[row][i] = true
			}
		}
	}
	return count
}
