// https://leetcode.com/problems/create-maximum-number/description/

package main

// pick up
// for 1 <= k <= m
// 1 <= i <= k, pick up i from nums1, and k-i from nums2
// choose the highest numbers from
// j := nums1[0:len(nums1)-i] where the largest number, nums1[j+1:len(nums1)-i+1]
// l := nums2[0:len(nums2)-k-i] where the largest number, nums2[l+1:len(nums1)-k-i+1]

// [3, 4, 6, 5], [9, 1, 2, 5, 8, 3], k + 5
// i = 1: [6], [9, 5, 8, 3]
// i = 2: [6, 5], [9, 8, 3]
// i = 3: [4, 6, 5], [9, 8]
// i = 4: [3, 4, 6, 5], [9]
// O(k*(m * n)): O(n+1)

// O(len*N)
func getSubArray1(array []int, arrayLen int) []int {
	if arrayLen == 0 {
		return nil
	}

	subArray := make([]int, arrayLen)
	startIndex := 0
	for count := 1; count <= arrayLen; count++ {
		maxIndex := startIndex
		for i := startIndex + 1; i < len(array)-arrayLen+count; i++ {
			if array[maxIndex] < array[i] {
				maxIndex = i
			}
		}
		subArray[count-1] = array[maxIndex]
		startIndex = maxIndex + 1
	}
	return subArray
}

// O(N)
// Using a stack to push an element and pops for older, smaller numbers always
func getSubArray2(array []int, arrayLen int) []int {
	// stack approach
	// we can drop the len(array)-arrayLen, like 2 elements to pick up 3 out of 5 array
	droppableCount := len(array) - arrayLen
	stack := make([]int, 0)
	for i := 0; i < len(array); i++ {
		for {
			if droppableCount <= 0 {
				break
			}
			if len(stack) == 0 {
				break
			}
			if stack[len(stack)-1] >= array[i] {
				break
			}
			stack = stack[:len(stack)-1]
			droppableCount--
		}
		stack = append(stack, array[i])
	}
	return stack[:arrayLen]
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	return maxNumber2(nums1, nums2, k)
}

func isGreater(array1 []int, array2 []int, array1Index int, array2Index int) bool {
	i := 0
	for ; array1Index+i < len(array1) && array2Index+i < len(array2) && array1[array1Index+i] == array2[array2Index+i]; i++ {
		// just add an index on an i to find the place
	}
	if array1Index+i >= len(array1) || array2Index+i >= len(array2) {
		// either of an array reaches to the end
		// if there is an element on the array1 still, then returns true so that it'll pick up a longer array
		// otherwise, returns false
		return array2Index+i >= len(array2)
	}

	return array1[array1Index+i] > array2[array2Index+i]
}

func merge(array1 []int, array2 []int) []int {
	result := make([]int, len(array1)+len(array2))
	array1Index := 0
	array2Index := 0
	for array1Index < len(array1) && array2Index < len(array2) {
		if isGreater(array1, array2, array1Index, array2Index) {
			result[array1Index+array2Index] = array1[array1Index]
			array1Index++
		} else {
			result[array1Index+array2Index] = array2[array2Index]
			array2Index++
		}
	}
	for ; array1Index < len(array1); array1Index++ {
		result[array1Index+array2Index] = array1[array1Index]
	}
	for ; array2Index < len(array2); array2Index++ {
		result[array1Index+array2Index] = array2[array2Index]
	}
	return result
}

func maxNumber1(nums1 []int, nums2 []int, k int) []int {
	getSubArray := getSubArray2

	// 6, 2, k = 5, then pick up at least 3 from nums1
	// initI := len(nums2) - k
	// if initI <= 0 {
	// 	initI = k - len(nums2)
	// }
	initI := 0

	result := make([]int, k)
	for subNums1Len := initI; subNums1Len <= k; subNums1Len++ {
		subNums2Len := k - subNums1Len
		if subNums1Len > len(nums1) || subNums2Len > len(nums2) {
			continue
		}
		if subNums2Len < 0 {
			// if the nums2 cannot create a subarray anymore, this no longer need to check
			break
		}

		subNums1 := getSubArray(nums1, subNums1Len)
		subNums2 := getSubArray(nums2, subNums2Len)
		merged := merge(subNums1, subNums2)
		if isGreater(merged, result, 0, 0) {
			result = merged
		}
		// fmt.Printf("%d: %+v, %+v, %+v, %+v\n", i, subNums1, subNums2, merged, result)
	}
	return result
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func maxNumber2(nums1 []int, nums2 []int, k int) []int {
	dp := func(array []int, k int) []int {
		if len(array) <= k {
			return array
		}
		i := 0
		for ; i+1 < len(array) && array[i] >= array[i+1]; i++ {
			// find an element which is a greater than the previous element
		}
		result := make([]int, 0)
		if i < len(array) {
			result = append(result,
				array[:i]...,
			)
			result = append(result,
				array[i+1:]...,
			)
		} else {
			result = append(result, array[:len(array)-1]...)
		}
		return result
	}

	dp1 := make([][]int, k+1)
	dp1[k] = getSubArray2(nums1, min(len(nums1), k))
	dp2 := make([][]int, k+1)
	dp2[k] = getSubArray2(nums2, min(len(nums2), k))
	for i := k - 1; i >= 0; i-- {
		dp1[i] = dp(dp1[i+1], i)
		dp2[i] = dp(dp2[i+1], i)

		// fmt.Printf("%d: %+v, %+v\n", i, dp1[i], dp2[i])
	}

	// 6, 2, k = 5, then pick up at least 3 from nums1
	initI := 0

	result := make([]int, k)
	for subNums1Len := initI; subNums1Len <= k; subNums1Len++ {
		subNums2Len := k - subNums1Len
		if subNums1Len > len(nums1) || subNums2Len > len(nums2) {
			continue
		}
		if subNums2Len < 0 {
			// if the nums2 cannot create a subarray anymore, this no longer need to check
			break
		}

		subNums1 := dp1[subNums1Len]
		subNums2 := dp2[subNums2Len]
		merged := merge(subNums1, subNums2)
		if isGreater(merged, result, 0, 0) {
			result = merged
		}

		// fmt.Printf("%d: %+v, %+v, %+v, %+v\n", subNums1Len, subNums1, subNums2, merged, result)
	}
	return result
}
