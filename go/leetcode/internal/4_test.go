package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	testCases := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			name:  "example 1",
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2,
		},
		{
			name:  "example 2",
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
		},
		{
			name:  "example 3",
			nums1: []int{1},
			want:  1,
		},
		{
			name:  "example 4",
			nums1: []int{1, 2, 3, 4},
			nums2: []int{5, 6, 7},
			want:  4,
		},
		{
			name:  "example 5",
			nums1: []int{1, 2, 3, 4},
			nums2: []int{1, 2, 3, 4, 4, 5, 5, 6, 7},
			want:  4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, findMedianSortedArrays1(tc.nums1, tc.nums2))
			assert.Equal(t, tc.want, findMedianSortedArrays2(tc.nums1, tc.nums2))
		})
	}
}
