package main

// Sort an array => find the value from 1
// Time: O(n log n)

// Cycle sorting
// Time complexity: O(n)
// Space complexity: O(1)
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		// only run once at most to put the values to indexes
		for {
			index := nums[i] - 1
			if index < 0 || len(nums) <= index {
				break
			}
			if nums[index] == nums[i] {
				break
			}

			nums[index], nums[i] = nums[i], nums[index]
		}
	}
	for i := 0; i < len(nums); i++ {
		value := nums[i]
		if value == i+1 {
			continue
		}
		return i + 1
	}
	return len(nums) + 1
}

// func abs(value int) int {
// 	if value < 0 {
// 		return -value
// 	}
// 	return value
// }

// Time complexity: O(n)
// Space complexity: O(1)
// func firstMissingPositive(nums []int) int {
// 	// In order to have a meaning for the negative value, remove values < 0
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] <= 0 {
// 			// The max nums length is 100k
// 			// Use math.MaxInt
// 			nums[i] = math.MaxInt
// 		}
// 	}

// 	// Update nums.
// 	// The original nums values are stored as abs(nums)
// 	// if nums[i] < 0, then the value i + 1 exists in the array nums
// 	// [3, 4, 0, 1], output: [-3, -4, 0, -1]
// 	for i := 0; i < len(nums); i++ {
// 		index := abs(nums[i]) - 1
// 		if index >= len(nums) {
// 			continue
// 		}
// 		if nums[index] < 0 {
// 			// this value is duplicated. No need to update
// 			continue
// 		}

// 		nums[index] = -1 * nums[index]
// 	}

// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] < 0 {
// 			continue
// 		}
// 		return i + 1
// 	}
// 	return len(nums) + 1
// }
