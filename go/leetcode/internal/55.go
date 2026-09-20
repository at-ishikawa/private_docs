package internal

func canJump(nums []int) bool {
	maxIndex := 0
	for i := 0; i < len(nums); i++ {
		if i > maxIndex {
			break
		}
		maxIndex = max(maxIndex, i+nums[i])
		if maxIndex >= len(nums)-1 {
			return true
		}
	}
	return false
}
