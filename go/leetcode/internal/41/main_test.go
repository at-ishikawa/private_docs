package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{1, 2, 0},
			want: 3,
		},
		{
			name: "example 2",
			nums: []int{3, 4, -1, 1},
			want: 2,
		},
		{
			name: "example 3",
			nums: []int{7, 8, 9, 11, 12},
			want: 1,
		},
		{
			name: "test case 1",
			nums: []int{1, 1},
			want: 2,
		},
		{
			name: "test case 2",
			nums: []int{0, 1, 2},
			want: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := firstMissingPositive(tc.nums)
			assert.Equal(t, tc.want, got)
		})
	}
}
