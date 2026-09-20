package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFinalState(t *testing.T) {
	testCases := []struct {
		name       string
		nums       []int
		k          int
		multiplier int
		want       []int
	}{
		{
			name: "example 1",
			nums: []int{
				2, 1, 3, 5, 6,
			},
			k:          5,
			multiplier: 2,
			want:       []int{8, 4, 6, 5, 6},
		},
		{
			name: "example 2",
			nums: []int{
				100000, 2000,
			},
			k:          2,
			multiplier: 1000000,
			want:       []int{999999307, 999999993},
		},
		{
			name:       "test case 1",
			nums:       []int{161209470},
			k:          56851412,
			multiplier: 39846,
		},
		{
			name:       "test case 2",
			nums:       []int{66307295, 441787703, 589039035, 322281864},
			k:          900900704,
			multiplier: 641725,
			want:       []int{664480092, 763599523, 886046925, 47878852},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := getFinalState(tc.nums, tc.k, tc.multiplier)
			assert.Equal(t, tc.want, got)
		})
	}
}
