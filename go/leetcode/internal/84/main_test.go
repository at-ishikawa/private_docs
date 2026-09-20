package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestRectangleArea(t *testing.T) {
	testCases := []struct {
		name    string
		heights []int
		want    int
	}{
		{
			name: "example 1",
			heights: []int{
				2, 1, 5, 6, 2, 3,
			},
			want: 10,
		},
		{
			name: "example 2",
			heights: []int{
				2, 4,
			},
			want: 4,
		},
		{
			name: "test case 1",
			heights: []int{
				4, 2, 0, 3, 2, 4, 3, 4,
			},
			want: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, largestRectangleArea1(tc.heights))
			assert.Equal(t, tc.want, largestRectangleArea2(tc.heights))
		})
	}
}
