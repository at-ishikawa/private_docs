// https://interviewing.io/questions/split-array-largest-sum

package main

import (
	"fmt"
	"math"
)

func main() {
	testCases := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "example 1",
			nums: []int{7, 2, 5, 10, 8},
			k:    3,
			want: 14,
		},
		{
			name: "example 2",
			nums: []int{1, 2, 3, 4, 5},
			k:    2,
			want: 9,
		},
		{
			name: "smallest",
			nums: []int{1},
			k:    1,
			want: 1,
		},
	}

	for _, tc := range testCases {
		fmt.Printf("Test %s: ", tc.name)
		got := minimizeSumOfSplittedArray2(tc.nums, tc.k)
		if tc.want != got {
			fmt.Printf("failed. Want %d, Got %d\n", tc.want, got)
		} else {
			fmt.Printf("passed\n")
		}
	}
}

func sum(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}

func maxArray(nums []int) int {
	result := math.MinInt
	for _, num := range nums {
		result = max(result, num)
	}
	return result
}

// Time complexity: O(N + N + N*log_2(SUM(nums)-MAX(nums))) => O(N*log_2(SUM(nums)))
// Space complexity: O(1)
func minimizeSumOfSplittedArray2(nums []int, k int) int {
	canSplit := func(nums []int, k int, mid int) bool {
		subArrayCount := 1
		currentSum := 0
		for i := 0; i < len(nums); i++ {
			if currentSum+nums[i] > mid {
				subArrayCount++
				currentSum = 0
			}
			currentSum += nums[i]
		}
		// fmt.Printf("%d, %d, %d\n", subArrayCount, k, currentSum)

		return subArrayCount <= k
	}

	left := maxArray(nums)
	right := sum(nums)
	for left <= right {
		mid := (left + right) / 2
		// fmt.Printf("[%d, %d]: %d\n", left, right, mid)

		if canSplit(nums, k, mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// Time Complexity: O(n^2)
// Space Complexity: O(n)
func minimizeSumOfSplittedArray1(nums []int, k int) int {
	// assume len(nums) < k

	// n-(k_1)
	// [5, 3, 6, 2, 10]
	// k =2
	// array 0: [5], array 1: [10]
	// array 0: [5, 3, 6], array 1: [10]
	// array 0: [5, 3, 6]: array 2: [2, 10]

	// k=3
	// array 0: [5], array last: [10]
	// array 0: [5],  arary 1: [3], array last: [10]

	// brute force
	// find all combinations of arrays
	//    recursively find array
	// calculate the maxSum => find min sum

	// generate where to divide
	// create k-1 array where 0 < divider_i <len(nums)-1
	// where divider_i means an array is divided by nums[:i], nums[i+1:]

	slices := make([][]int, 0)
	return divideArray(nums, k, 0, 0, []int{}, 0, slices)
}

func divideArray(nums []int, k int, numIndex int, maxSum int, subArray []int, currentArraySum int, slices [][]int) int {
	if k <= 1 {
		for i := numIndex; i < len(nums); i++ {
			currentArraySum += nums[i]
		}
		// fmt.Printf("%d, %d, %#v\n", maxSum, currentArraySum, slices)
		// fmt.Printf("%d, %d, %d, %d\n", k, numIndex, maxSum, currentArraySum)
		return max(maxSum, currentArraySum)
	}
	if numIndex >= len(nums) {
		return math.MaxInt
	}
	if numIndex+k > len(nums) {
		return math.MaxInt
	}

	minSum := math.MaxInt
	subArrayLength := len(subArray)
	for i := numIndex; i < len(nums); i++ {
		subArray = append(subArray, nums[i])
		currentArraySum += nums[i]

		// fmt.Printf("%s: (%d, %d, %d)\n", strings.Repeat(" ", k*2), nums[i], maxSum, currentArraySum)
		minSum = min(
			minSum,
			// in a case if an array is split
			divideArray(nums, k-1, i+1, max(maxSum, currentArraySum), []int{}, 0,
				append(slices, subArray),
			),
			// in a case if an array hasn't been split
			divideArray(nums, k, i+1, maxSum, subArray, currentArraySum, slices),
		)
	}
	if numIndex > 0 {
		subArray = subArray[:subArrayLength]
	}

	return minSum
}
