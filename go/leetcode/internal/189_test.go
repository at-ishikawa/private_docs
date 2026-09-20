package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	testCases := map[string]struct {
		nums []int
		k    int
		want []int
	}{
		"example 1": {
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		"example 2": {
			nums: []int{1, -100, 3, 99},
			k:    2,
			want: []int{3, 99, 1, -100},
		},
		"submit case 1": {
			nums: []int{-1},
			k:    2,
			want: []int{-1},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Run("rotate 1", func(t *testing.T) {
				nums := make([]int, len(tc.nums))
				copy(nums, tc.nums)
				rotate1(nums, tc.k)
				got := nums
				assert.Equal(t, tc.want, got)
			})

			t.Run("rotate 2", func(t *testing.T) {
				nums := make([]int, len(tc.nums))
				copy(nums, tc.nums)
				rotate2(nums, tc.k)
				got := nums
				assert.Equal(t, tc.want, got)
			})

			t.Run("rotate 3", func(t *testing.T) {
				nums := make([]int, len(tc.nums))
				copy(nums, tc.nums)
				rotate3(nums, tc.k)
				got := nums
				assert.Equal(t, tc.want, got)
			})
		})
	}
}
