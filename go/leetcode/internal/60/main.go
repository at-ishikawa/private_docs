package main

import (
	"slices"
	"strconv"
)

// 123
// 132
// 213
// 231
// 312
// 321
// There are 3*2*1 => 6.
// The first digit: k / 3: 0 => 1, 1 => 2, 2 => 3
// The second digit: (k - (k / 3)*3) / 2: 0 => smaller, 1 => bigger

// 4! => 24
// 3421: 17

// 12: 3124
// 15: 3241
// 16: 3412
// 17: 3421
// 18: 4123

// There are 6 permutations for each first digit. To figure out, divide by 6 to figure out which one.
//     6 is from 24 (all) / 4 (n)
// There are 2 permutations for each second digit in these 6 permutations, divided by 2 to figure out which one
//     2 is from 6 (remaining) / 3 (n-1)
// There are

// First: k / 6 => 3, for next k - 12: 5
// Second: 5 / 2 => 2, use 4. k - 4 = 1
// third: 1 / 1 => 1, use 2,
// last: use 1

// Time: O(n) Space O(1)
func getPermutation(n int, k int) string {
	allCount := 1
	digits := make([]int, 0)
	for i := 1; i <= n; i++ {
		allCount *= i
		digits = append(digits, i)
	}

	remainingCombinations := k - 1
	combinations := allCount
	permutation := ""
	for i := 0; i < n; i++ {
		//
		// 3: 213
		// 4: 231
		//
		// 2 / 2 => 1, next 1
		// 1 / 1 => 1, next
		//
		combinations = combinations / (n - i)
		digit := (remainingCombinations / combinations) + 1

		// fmt.Printf("%d/%d, %d, %s, %#v\n", remainingCombinations, combinations, digit, permutation, digits)

		permutation = permutation + strconv.Itoa(digits[digit-1])
		digits = slices.Delete(digits, digit-1, digit)

		// 17-12
		remainingCombinations = remainingCombinations - (remainingCombinations/combinations)*combinations
	}
	return permutation
}
