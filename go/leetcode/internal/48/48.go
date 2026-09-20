package internal48

import "slices"

// 1, 2, 3, 4, 5
// 6, 7, 8, 9, 10
// 11, 12, 13, 14, 15
// 16, 17, 18, 19, 20
// 21, 22, 23, 24, 25

// 21, 16, 11, 6, 1
// 22, 17, 12, 7, 2
// 23, 18, 13, 8, 3
// 24, 19, 14, 0, 4
// 25, 20, 15, 10, 5

func rotate(matrix [][]int) {
	rotate2(matrix)
}

func rotate2(matrix [][]int) {
	slices.Reverse(matrix)
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func rotate1(matrix [][]int) {
	n := len(matrix)
	// move cell[i][j] to a counter clockwise 3 times
	// move
	//   1. swap (cell[i][j], cell[n-j][i])
	//   2. swap (cell[n-j][i], cell[n-i][n-j])
	//   3. swap (cell[n-i][n-j], cell[j][n-i])
	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ {
			nj := n - j - 1
			ni := n - i - 1

			indexes := [][]int{
				{i, j},
				{nj, i},
				{ni, nj},
				{j, ni},
			}
			for k := 1; k < len(indexes); k++ {
				srcX, srcY := indexes[k-1][0], indexes[k-1][1]
				dstX, dstY := indexes[k][0], indexes[k][1]
				matrix[srcX][srcY], matrix[dstX][dstY] = matrix[dstX][dstY], matrix[srcX][srcY]
			}
			// matrix[i][j], matrix[nj][i] = matrix[nj][i], matrix[i][j]
			// matrix[nj][i], matrix[ni][nj] = matrix[ni][nj], matrix[nj][i]
			// matrix[ni][nj], matrix[j][ni] = matrix[j][ni], matrix[ni][nj]
		}
	}
}
