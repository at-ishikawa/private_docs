package main

// max area for width * height where
// width = (j - i + 1) where i < j, and height = min(heights[k]) where i <= k <= j
func largestRectangleArea(heights []int) int {
	return largestRectangleArea1(heights)
}

// find from the point i, calculate the area to the j, and find the max area

// Calculate the left position and right position on each i where 0 <= i < len(heights)
// Use th previously calculated values on lefts and rights so that the calculations are faster than O(N^2)

// The worst case: [1, 2, 3, 4, 5]
// lefts: [1, 2, 3, 4, 5], rights: [5, 5, 5, 5, 5] (rights can be calculated only one lookup on the rights array, so O(N))
func largestRectangleArea1(heights []int) int {
	lefts := make([]int, len(heights))
	lefts[0] = 0
	for index := 1; index < len(heights); index++ {
		left := index - 1
		for ; left >= 0 && heights[left] >= heights[index]; left = lefts[left] - 1 {
		}
		// fmt.Printf("%d, %d, %d\n", index, heights[index], left)
		lefts[index] = left + 1
	}
	rights := make([]int, len(heights))
	rights[len(heights)-1] = len(heights) - 1
	for index := len(heights) - 2; index >= 0; index-- {
		right := index + 1
		for ; right < len(heights) && heights[index] <= heights[right]; right = rights[right] + 1 {
		}
		// fmt.Printf("%d, %d, %d\n", index, heights[index], right)
		rights[index] = right - 1
	}

	max := 0
	for i := 0; i < len(heights); i++ {
		area := heights[i] * (rights[i] - lefts[i] + 1)
		if max < area {
			max = area
		}
	}
	return max
}

// stack approach: https://leetcode.com/problems/largest-rectangle-in-histogram/solutions/28900/short-and-clean-o-n-stack-based-java-solution/

// [2, 1, 5, 6, 2, 3]
// 2 => index stack: [0]
// [2, 1] => index stack: [], area
func largestRectangleArea2(heights []int) int {
	n := len(heights)
	indexStack := make([]int, 0)

	// the area should be the heights[index] * (right - left) where right = index, left = the previous index of a stack
	// The left is the previous index of a stack because the value of a height is bigger than the height of previous element in a stack
	getArea := func(right int) int {
		index := indexStack[len(indexStack)-1]
		indexStack = indexStack[:len(indexStack)-1]

		left := 0
		if len(indexStack) > 0 {
			left = indexStack[len(indexStack)-1] + 1
		}
		width := right - left
		height := heights[index]
		return height * width
	}

	maxArea := 0
	for index := 0; index <= n; {
		var height int
		if index < n {
			height = heights[index]
		}

		// Unless a bar in the stack is the largest bar, push the current bar into the stack
		if len(indexStack) == 0 || height >= heights[indexStack[len(indexStack)-1]] {
			indexStack = append(indexStack, index)
			index++
			continue
		}

		// the stack has indices for heights where values are monotonically increasing
		// if a current bar is shorter, then the areas from the previous bars in the stack can be calculated
		area := getArea(index)
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
