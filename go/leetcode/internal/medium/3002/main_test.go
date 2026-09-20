package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumSetSize(t *testing.T) {
	testCases := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  int
	}{
		{
			name:  "example 1",
			nums1: []int{1, 2, 1, 2},
			nums2: []int{1, 1, 1, 1},
			want:  2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, maximumSetSize(tc.nums1, tc.nums2))
		})
	}
}
