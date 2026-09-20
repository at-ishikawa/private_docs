package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumRemovables(t *testing.T) {
	testCases := []struct {
		name      string
		s         string
		p         string
		removable []int
		want      int
	}{
		{
			name:      "example 1",
			s:         "abcacb",
			p:         "ab",
			removable: []int{3, 1, 0},
			want:      2,
		},
		{
			name:      "example 2",
			s:         "abcbddddd",
			p:         "abcd",
			removable: []int{3, 2, 1, 4, 5, 6},
			want:      1,
		},
		{
			name:      "example 3",
			s:         "abcab",
			p:         "abc",
			removable: []int{0, 1, 2, 3, 4},
			want:      0,
		},
		{
			name:      "test case 1",
			s:         "qlevcvgzfpryiqlwy",
			p:         "qlecfqlw",
			removable: []int{12, 5},
			want:      2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, maximumRemovals(tc.s, tc.p, tc.removable))
		})
	}
}
