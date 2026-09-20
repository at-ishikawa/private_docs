package internal

import "math"

// size: m and size: n
// median merged_nums1_nums2(m + n) / 2
// Time complexity: O(log(m + n))
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArrays2(nums1, nums2)
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays2(nums2, nums1)
	}

	low := 0
	high := len(nums1)
	lower := (len(nums1) + len(nums2) + 1) / 2
	for low <= high {
		num1Index := (low + high) / 2
		num2Index := lower - num1Index

		lowNum1 := math.MinInt
		lowNum2 := math.MinInt
		highNum1 := math.MaxInt
		highNum2 := math.MaxInt
		if num1Index >= 0 && len(nums1) > num1Index {
			highNum1 = nums1[num1Index]
		}
		if num2Index >= 0 && len(nums2) > num2Index {
			highNum2 = nums2[num2Index]
		}
		if num1Index-1 >= 0 && len(nums1) > num1Index-1 {
			lowNum1 = nums1[num1Index-1]
		}
		if num2Index-1 >= 0 && len(nums2) > num2Index-1 {
			lowNum2 = nums2[num2Index-1]
		}
		if lowNum1 <= highNum2 && lowNum2 <= highNum1 {
			if (len(nums1)+len(nums2))%2 == 1 {
				return float64(max(lowNum1, lowNum2))
			}
			return float64(max(lowNum1, lowNum2)+min(highNum1, highNum2)) / 2
		}

		if highNum2 < lowNum1 {
			high = num1Index - 1
		} else {
			low = num1Index + 1
		}
	}
	return 0
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	num1Index := len(nums1) - 1
	num2Index := len(nums2) - 1
	previousValue := math.MaxInt
	var median int

	// go down to the number to reach to the index
	medianIndex := (len(nums1) + len(nums2)) / 2
	for count := 0; count <= medianIndex; count++ {
		previousValue = median
		if len(nums1) == 0 || num1Index < 0 {
			median = nums2[num2Index]
			num2Index--
			continue
		}
		if len(nums2) == 0 || num2Index < 0 {
			median = nums1[num1Index]
			num1Index--
			continue
		}

		if nums1[num1Index] < nums2[num2Index] {
			median = nums2[num2Index]
			num2Index--
		} else {
			median = nums1[num1Index]
			num1Index--
		}
	}

	if (len(nums1)+len(nums2))%2 == 1 {
		return float64(median)
	}
	return float64(median+previousValue) / 2
}
