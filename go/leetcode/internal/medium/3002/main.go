// https://leetcode.com/problems/maximum-size-of-a-set-after-removals/

package main

// nums1 = [1,2,3,4,5,6], nums2 = [2,3,2,3,2,3]
// num1: 1-6: 1, nums2: [2:3, 3:2]
// - We want to remove duplicated elements on each num, as well as both nums
// - After removing duplicates, there are still k1 and k2 elements need to be removed in nums1 and nums2
// - n - k1 - k2 => final result
// Q: if there is only one element on each nums, which one we should choose?
//	[1, 1, 2, 3], [1, 4, 5, 6], [1, 2], [1, 4]

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](slice []T) Set[T] {
	result := make(Set[T], 0)
	for _, val := range slice {
		result[val] = struct{}{}
	}
	return result
}

func IntersectSets[T comparable](set1 Set[T], set2 Set[T]) Set[T] {
	result := make(Set[T], 0)
	for key := range set1 {
		if _, ok := set2[key]; !ok {
			continue
		}
		result[key] = struct{}{}
	}
	return result
}

func maximumSetSize(nums1 []int, nums2 []int) int {
	set1 := NewSet(nums1)
	set2 := NewSet(nums2)
	intersected := IntersectSets(set1, set2)

	s1 := len(set1)
	s2 := len(set2)

	// [1, 2, 1, 2] vs [1, 2],
	ans := min(s1, len(nums1)/2) + min(s2, len(nums2)/2)
	return min(ans, s1+s2-len(intersected))
}
