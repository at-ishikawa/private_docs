package internal

func rotate(nums []int, k int) {
	rotate3(nums, k)
}

// Time: O(N*k)
// Space: O(1)
// Time limit exceeded
func rotate1(nums []int, k int) {
	k %= len(nums)
	for i := 0; i < k; i++ {
		tmp := nums[len(nums)-1]
		for endIndex := len(nums) - 2; endIndex >= 0; endIndex-- {
			nums[endIndex+1] = nums[endIndex]
		}
		nums[0] = tmp
	}
}

// Time cmoplexity: O(n)
// Space complexity: O(n)
func rotate2(nums []int, k int) {
	k %= len(nums)

	endIndex := len(nums) - k
	newSlice := append(nums[endIndex:], nums[:endIndex]...)
	copy(nums, newSlice)
}

// Time complexity: O(N + k + N-k)=O(N)
// Space complexity: O(1)
func rotate3(nums []int, k int) {
	k %= len(nums)

	reverse := func(slice []int, start, end int) {
		for start < end {
			tmp := slice[start]
			slice[start] = slice[end]
			slice[end] = tmp
			start++
			end--
		}
	}

	endIndex := len(nums) - 1
	reverse(nums, 0, endIndex)
	reverse(nums, 0, k-1)
	reverse(nums, k, endIndex)
}
