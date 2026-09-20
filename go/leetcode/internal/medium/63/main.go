// https://leetcode.com/problems/unique-paths-ii/description/
package main

// [0, 0] => [0,1]: only one path
// [0, 0] => [1,0]: only one path
// [0, 1], [1,0] => [1, 1]: 2 paths
// [y-1, x] + [y, x-1] => [y, x]
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	height := len(obstacleGrid)
	width := len(obstacleGrid[0])

	dp := make([][]int, height)
	for i := 0; i < height; i++ {
		dp[i] = make([]int, width)
	}

	for i := 0; i < width; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}
	for y := 1; y < height; y++ {
		if obstacleGrid[y][0] != 1 {
			dp[y][0] = dp[y-1][0]
		}
		for x := 1; x < width; x++ {
			if obstacleGrid[y][x] == 1 {
				continue
			}
			dp[y][x] = dp[y-1][x] + dp[y][x-1]
		}
	}

	return dp[height-1][width-1]
}
