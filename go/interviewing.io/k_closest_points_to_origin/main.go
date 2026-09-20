/**
Given a list of tuples that represent (X, Y) coordinates on an XY plane and an integer K, return a list of the K-closest points to the origin (0, 0).

Example Inputs and Outputs
Example 1
- Input:
    -  points = [[5, 5], [3, 3], [4, 4]], k = 2
- Output:
    - [[3, 3], [4, 4]] or [[4, 4], [3, 3]]

Example 2
- Input:
    - points = [[-1, 4], [5, 3], [-1, -1], [8, -6], [1, 2]], k = 2
- Output:
    - [[-1, -1], [1, 2]] or [[1, 2], [-1, -1]]

Constraints
- The number of nodes in the list is in the range [0, 5000]
- K is >= 0 and <= the length of the input list
*/

package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	testCases := []struct {
		name   string
		tuples [][]int
		k      int
		want   [][]int
	}{
		{
			name: "example 1",
			tuples: [][]int{
				{5, 5},
				{3, 3},
				{4, 4},
			},
			k: 2,
			want: [][]int{
				{3, 3},
				{4, 4},
			},
		},
		{
			name: "example 2",
			tuples: [][]int{
				{-1, 4},
				{5, 3},
				{-1, -1},
				{8, -6},
				{1, 2},
			},
			k: 2,
			want: [][]int{
				{-1, -1},
				{1, 2},
			},
		},
	}

	for _, tc := range testCases {
		got := kClosestPoints(tc.tuples, tc.k)
		fmt.Printf("Test %s: ", tc.name)
		if reflect.DeepEqual(tc.want, got) {
			fmt.Printf("Passed.")
		} else {
			fmt.Printf("Failed. Want: %#v, Got: %#v", tc.want, got)
		}
		fmt.Println()
	}
}

func getDistance(tuple []int) int {
	return tuple[0]*tuple[0] + tuple[1]*tuple[1]
}

func kClosestPoints(tuples [][]int, k int) [][]int {
	sort.Slice(tuples, func(i, j int) bool {
		point1 := tuples[i]
		point2 := tuples[j]
		return getDistance(point1) < getDistance(point2)
	})
	result := make([][]int, k)
	for i := 0; i < k; i++ {
		result[i] = tuples[i]
	}
	return result
}
