// https://leetcode.com/problems/minimum-moves-to-capture-the-queen/description/
package main

import "math"

func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
	return minMovesToCaptureTheQueen1(a, b, c, d, e, f)
}

const (
	rook   = 1
	bishop = 2
	queen  = 9
)

type Backtracker struct {
	rookDirections   [][]int
	bishopDirections [][]int
}

func (tracker Backtracker) backtrack(
	chessboard [9][9]int,
	rookCell []int,
	bishopCell []int,
	steps int,
) int {
	if steps >= 2 {
		return math.MaxInt
	}

	result := math.MaxInt
	for index, directions := range [][][]int{tracker.rookDirections, tracker.bishopDirections} {
		var y, x int
		var pawn int
		if index == 0 {
			// rook
			pawn = rook
			y, x = rookCell[0], rookCell[1]
		} else {
			pawn = bishop
			y, x = bishopCell[0], bishopCell[1]
		}

		chessboard[y][x] = 0
		for _, direction := range directions {
			for i := 0; i < 8; i++ {
				nextY, nextX := direction[0]*i+y, direction[1]*i+x
				if nextY <= 0 || nextX <= 0 ||
					nextY >= len(chessboard) || nextX >= len(chessboard[nextY]) {
					break
				}

				if chessboard[nextY][nextX] == queen {
					chessboard[y][x] = pawn
					return steps + 1
				}
				if chessboard[nextY][nextX] != 0 {
					// if there is a bishop, no more moving to the direction
					break
				}

				nextRookCell := rookCell
				nextBishopCell := bishopCell
				if pawn == rook {
					nextRookCell = []int{nextY, nextX}
				} else {
					nextBishopCell = []int{nextY, nextX}
				}
				chessboard[nextY][nextX] = pawn
				result = min(result, tracker.backtrack(
					chessboard,
					nextRookCell,
					nextBishopCell,
					steps+1,
				))
				chessboard[nextY][nextX] = 0
			}
		}
		chessboard[y][x] = pawn
	}
	return result
}

func minMovesToCaptureTheQueen2(a int, b int, c int, d int, e int, f int) int {
	chessboard := [9][9]int{}
	chessboard[a][b] = 1     // rook
	chessboard[c][d] = 2     // bishop
	chessboard[e][f] = queen // queen

	return Backtracker{
		rookDirections: [][]int{
			{1, 0},
			{-1, 0},
			{0, 1},
			{0, -1},
		},
		bishopDirections: [][]int{
			{1, 1},
			{-1, 1},
			{1, -1},
			{-1, -1},
		},
	}.backtrack(chessboard,
		[]int{a, b},
		[]int{c, d},
		0,
	)
}

// Time complexity: O(N) (N=chessboard size)
// Space complexity: O(1)
// answer:
// 1 if a bishop can be get by a rook or queen's first move
// 2 otherwise
//
// how to check a queen is in the same row/column of a rook, or a diagnoally reached from the bisop?
// rook => queen: a = e, or b = f (either horizontally or vertically matching places)
// bishop => queen: c + i = d and e + i = f where <= i <= 8
// edge case: bishop is in the way of a rook.
//   - a = c = e => b < d < f || b > d > f
//   - b = d = f => a < c < e || a > c > e
type Cell struct {
	y int
	x int
}

func minMovesToCaptureTheQueen1(a int, b int, c int, d int, e int, f int) int {

	bishop := Cell{
		y: c,
		x: d,
	}
	queen := Cell{
		y: e,
		x: f,
	}
	rook := Cell{
		y: a,
		x: b,
	}

	// bishop can take a queen!
	directions := [][]int{
		{1, 1},
		{-1, 1},
		{1, -1},
		{-1, -1},
	}
	for _, direction := range directions {
		for i := 0; i <= 8; i++ {
			yDelta, xDelta := direction[0]*i, direction[1]*i
			y := bishop.y + yDelta
			x := bishop.x + xDelta
			if y == rook.y && x == rook.x {
				// if there is a rook in the middle of the way, cannot take a bishop
				break
			}
			if y == queen.y && x == queen.x {
				return 1
			}
		}
	}

	if rook.y == queen.y {
		if rook.y == bishop.y {
			if rook.x < bishop.x && bishop.x < queen.x {
				return 2
			}
			if rook.x > bishop.x && bishop.x > queen.x {
				return 2
			}
		}
		return 1
	}
	if rook.x == queen.x {
		if rook.x == bishop.x {
			if rook.y < bishop.y && bishop.y < queen.y {
				return 2
			}
			if rook.y > bishop.y && bishop.y > queen.y {
				return 2
			}
		}
		return 1
	}
	return 2
}
