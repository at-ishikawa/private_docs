// https://leetcode.com/problems/remove-k-digits/description/
package main

import (
	"strings"
)

// 159, k=1
// 15, not 59, nor 19
// 1591, k=1
// 151, 159, 191
// 1579, k=1
// 157, 159, 179
// 1528, k=1
// 152, 158, 128
// 5283, k=1
// 528, 523, 283
// 5923, k=2
// 52, 53, 92, 93, 23
// 3958, k=2
// 39, 35, 38, 95, 98, 58

func removeKdigits(num string, k int) string {
	stack := make([]int, 0)
	numIndex := 0
	for ; numIndex < len(num); numIndex++ {
		digit := int(num[numIndex] - '0')
		if len(stack) > 0 {
			for 0 < k && 0 < len(stack) {
				top := stack[len(stack)-1]
				if top <= digit {
					break
				}

				// remove the top from a stack
				stack = stack[:len(stack)-1]
				k--
			}
			if k == 0 {
				break
			}
		}
		stack = append(stack, digit)
	}
	// remove remaining element from stack as long as there are k removals
	if k > 0 && len(stack) > 0 {
		if len(stack)-k > 0 {
			stack = stack[:len(stack)-k]
		} else {
			stack = nil
		}
	}

	builder := strings.Builder{}
	for i := 0; i < len(stack); i++ {
		if builder.Len() == 0 && stack[i] == 0 {
			continue
		}
		builder.WriteByte(byte(stack[i] + '0'))
	}
	for ; numIndex < len(num); numIndex++ {
		if builder.Len() == 0 && num[numIndex] == '0' {
			continue
		}
		builder.WriteByte(num[numIndex])
	}
	if builder.Len() == 0 {
		return "0"
	}

	return builder.String()
}
