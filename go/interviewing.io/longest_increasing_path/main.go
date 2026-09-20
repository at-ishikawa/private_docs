package main

import "fmt"

// Time complexity: O(n*m)
// Space complexity: O(n*m)

func dfs(matrix [][]int, maxPaths [][]int, x, y int, minCell int) int {
	if x < 0 || y < 0 || y >= len(matrix) || x >= len(matrix[0]) {
		return 0
	}
	if minCell >= matrix[y][x] {
		return 0
	}
	if maxPaths[y][x] > 0 {
		return maxPaths[y][x]
	}

	result := 0
	nextSteps := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	for _, nextStep := range nextSteps {
		nextY := y + nextStep[0]
		nextX := x + nextStep[1]
		length := dfs(matrix, maxPaths, nextX, nextY, matrix[y][x])
		if result < length {
			result = length
		}
	}
	result++
	maxPaths[y][x] = result
	return result
}

func longestIncreasingPath(matrix [][]int) int {
	maxPaths := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		maxPaths[i] = make([]int, len(matrix[i]))
	}

	result := 0
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			length := dfs(matrix, maxPaths, y, x, 0)
			if result < length {
				result = length
			}
		}
	}
	return result
}

func main() {
	testCases := []struct {
		name   string
		matrix [][]int
		want   int
	}{
		{
			name: "example 1",
			matrix: [][]int{
				{9, 9, 4},
				{6, 6, 8},
				{2, 1, 1},
			},
			want: 4,
		},
		{
			name: "example 2",
			matrix: [][]int{
				{3, 4, 5},
				{3, 2, 6},
				{2, 2, 1},
			},
			want: 4,
		},
	}
	for _, tc := range testCases {
		got := longestIncreasingPath(tc.matrix)
		fmt.Printf("Test: %s: Want: %d, Got: %d\n", tc.name, tc.want, got)
	}
}
