// https://leetcode.com/problems/find-kth-bit-in-nth-binary-string/description/

package main

import "strings"

func findKthBit(n int, k int) byte {
	k--
	s := strings.Builder{}
	s.WriteByte('0')
	for i := 2; i <= n; i++ {
		previous := s.String()
		s.WriteByte('1')
		for i := len(previous) - 1; i >= 0; i-- {
			if previous[i] == '0' {
				s.WriteByte('1')
			} else {
				s.WriteByte('0')
			}
		}
		if s.Len() > k {
			return s.String()[k]
		}
	}
	return s.String()[k]
}
