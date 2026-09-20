// https://leetcode.com/problems/select-cells-in-grid-with-maximum-score/description/
package main

import (
	"sort"
)

// check all combinations by recursive function on each row
// options: max values, unique values

// From the maximum values, try to pick up the right one from the next possible values

// [8, 7, 5, 1, 1] => 5
// [8, 1, 1, 1, 1] => 8
// [8, 7, 1, 1, 1] => 7
// [7, 5, 2, 1, 1] => 2

// [8, 7, 5, 1, 1] => 5
// [8, 7, 1, 1, 1] => 8
// [8, 7, 6, 1, 1] => 6
// [7, 5, 2, 1, 1] => 7,

// Find 8 and check next values for 8 in 3
// all of them are 7, so find next values for 7, then find 5 in 2 rows, 1, 6

func maxCell(array []int, usedValues map[int]int) int {
	result := -1
	for _, v := range array {
		if usedValues[v] > 0 {
			continue
		}

		if result < v {
			result = v
		}
	}
	return result
}

func getSumsFromValueMap(m map[int]int) int {
	sum := 0
	for v := range m {
		sum += v
	}
	return sum
}

func getMaxSum(grid [][]int, usedRows map[int]int, usedValues map[int]int, _ int) int {
	// fmt.Printf("%+v, %+v\n", usedRows, usedValues)
	if len(usedRows) == len(grid) {
		// we found all values that are used from possible values
		getSumsFromValueMap(usedValues)
	}

	maxValue := -1
	var maxValueRowIndexes []int
	for rowIndex, row := range grid {
		if usedRows[rowIndex] > 0 {
			continue
		}

		max := maxCell(row, usedValues)
		if maxValue < max {
			maxValue = max
			maxValueRowIndexes = []int{
				rowIndex,
			}
		} else if maxValue == max {
			maxValueRowIndexes = append(maxValueRowIndexes, rowIndex)
		}
	}

	if maxValue == -1 {
		// if no value cannot be found, for example, no more unique value in a row
		// then return the sum to a user
		return getSumsFromValueMap(usedValues)
	}
	if len(maxValueRowIndexes) < 2 {
		// max value is unique, we can decide to pick up this value and move to next value
		rowIndex := maxValueRowIndexes[0]
		usedRows[rowIndex]++
		usedValues[maxValue]++
		result := getMaxSum(grid, usedRows, usedValues, maxValue)
		delete(usedRows, rowIndex)
		delete(usedValues, maxValue)
		return result
	}

	// if there are multiple candidates, then we need to find values based on next values
	// at first, brute force
	result := 0
	for _, rowIndex := range maxValueRowIndexes {
		max := maxCell(grid[rowIndex], usedValues)
		if max != maxValue {
			panic(rowIndex)
		}

		// todo: usedValues should not be updated all the time, but instead before for loop
		usedValues[maxValue]++
		usedRows[rowIndex]++
		candidate := getMaxSum(grid, usedRows, usedValues, maxValue)
		delete(usedRows, rowIndex)
		usedValues[maxValue] = 0
		delete(usedValues, maxValue)
		if candidate > result {
			result = candidate
		}
	}

	return result
}

func maxScore(grid [][]int) int {
	return maxScore2(grid)
}

func maxScore1(grid [][]int) int {
	for index := range grid {
		sort.Slice(grid[index], func(i, j int) bool {
			return grid[index][i] > grid[index][j]
		})
	}

	usedRows := make(map[int]int, 0)
	usedValues := make(map[int]int, 0)
	return getMaxSum(grid, usedRows, usedValues, 0)
}

type Cell struct {
	row    int
	column int
	value  int
}

func maxScore2(grid [][]int) int {
	// O(N*M*log(NM))
	flatten := make([]Cell, len(grid)*len(grid[0]))
	for row := range grid {
		for column := range grid[row] {
			flatten[row*len(grid[0])+column] = Cell{
				value:  grid[row][column],
				row:    row,
				column: column,
			}
		}
	}
	sort.Slice(flatten, func(i, j int) bool {
		return flatten[i].value > flatten[j].value
	})

	usedRows := make(map[int]int, 0)
	caches := make(map[int]int, 0)
	return solve(flatten, 0, len(grid), usedRows, caches)
}

func getRowCacheKey(cellIndex int, usedRows map[int]int) int {
	// THe length of rows is at most 10, so use an index on the bit > 10
	cacheKey := cellIndex << 11
	for row := range usedRows {
		cacheKey |= 1 << row
	}
	return cacheKey
}

// N=len(grid), M=len(grid[0])
// Time complexity: O(N*M*M^N) = O(M^N)
//  1. a for loop without recursive : N*M at the worst case (all values are the same)
//  2. a recursive: M^N. There are M choices on each row at the worst case scenario.
//
// Space complexity:
//  1. usedRows: O(N)
//  2. caches: cellIndex: M*N, row combinations: M^N
func solve(cells []Cell, cellIndex int, rowCount int, usedRows map[int]int, caches map[int]int) int {
	cacheKey := getRowCacheKey(cellIndex, usedRows)
	if cache, ok := caches[cacheKey]; ok {
		return cache
	}

	if len(usedRows) >= rowCount {
		return 0
	}
	if cellIndex >= len(cells) {
		return 0
	}

	nextValueCellIndex := cellIndex
	for ; nextValueCellIndex < len(cells) && cells[cellIndex].value == cells[nextValueCellIndex].value; nextValueCellIndex++ {
	}

	result := 0
	// pick up the same value in the other row and pick up the largest value
	for i := cellIndex; i < nextValueCellIndex; i++ {
		row := cells[i].row
		if usedRows[row] > 0 {
			continue
		}

		usedRows[row]++
		candidate := cells[i].value + solve(cells, nextValueCellIndex, rowCount, usedRows, caches)
		delete(usedRows, row)
		if result < candidate {
			result = candidate
		}
	}
	if result == 0 {
		// if the same row with the value was used already, try to find the different value
		result = solve(cells, nextValueCellIndex, rowCount, usedRows, caches)
	}

	caches[cacheKey] = result
	return result
}
