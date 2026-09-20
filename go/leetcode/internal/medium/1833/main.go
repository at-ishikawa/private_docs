// https://leetcode.com/problems/maximum-ice-cream-bars/description/
// Counting sort:
//
//	https://www.geeksforgeeks.org/counting-sort/
package main

import (
	"slices"
)

// Time complexity: O(Max(Costs)+3N)
// Space complexity: O(Max(costs))
func maxIceCream(costs []int, coins int) int {
	sortedCosts := countSort(costs)

	countIceCream := 0
	for _, cost := range sortedCosts {
		if coins < cost {
			break
		}
		countIceCream++
		coins -= cost
	}
	return countIceCream
}

func countSort(array []int) []int {
	counts := make([]int, slices.Max(array)+1)
	for _, value := range array {
		counts[value]++
	}
	// Update counts to a cumulative array
	for i := 1; i < len(counts); i++ {
		counts[i] = counts[i] + counts[i-1]
	}

	// Update an array in-place
	sortedArray := make([]int, len(array))
	for i := 0; i < len(array); i++ {
		value := array[i]
		sortedArray[counts[value]-1] = value
		counts[value]--
	}
	return sortedArray
}
