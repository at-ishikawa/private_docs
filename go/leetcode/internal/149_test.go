package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxPoints(t *testing.T) {
	testCases := []struct {
		name   string
		points [][]int
		want   int
	}{
		{
			name: "example 1",
			points: [][]int{
				{1, 1},
				{2, 2},
				{3, 3},
			},
			want: 3,
		},
		{
			name: "example 2",
			points: [][]int{
				{1, 1},
				{3, 2},
				{5, 3},
				{4, 1},
				{2, 3},
				{1, 4},
			},
			want: 4,
		},
		{
			name: "example 3: x=4",
			points: [][]int{
				{4, 5},
				{4, -1},
				{4, 0},
			},
			want: 3,
		},
		{
			name: "example 4: y=4",
			points: [][]int{
				{5, 4},
				{-1, 4},
				{0, 4},
			},
			want: 3,
		},
		{
			name: "failed",
			points: [][]int{
				{54, 153},
				{1, 3},
				{0, -72},
				{-3, 3},
				{12, -22},
				{-60, -322},
				{0, -5},
				{-5, 1},
				{5, 5},
				{36, 78},
				{3, -4},
				{5, 0},
				{0, 4},
				{-30, -197},
				{-5, 0},
				{60, 178},
				{0, 0},
				{30, 53},
				{24, 28},
				{4, 5},
				{2, -2},
				{-18, -147},
				{-24, -172},
				{-36, -222},
				{-42, -247},
				{2, 3},
				{-12, -122},
				{-54, -297},
				{6, -47},
				{-5, 3},
				{-48, -272},
				{-4, -2},
				{3, -2},
				{0, 2},
				{48, 128},
				{4, 3},
				{2, 4},
			},
			want: 18,
		},
		{
			name: "failed 2",
			points: [][]int{
				{5151, 5150},
				{0, 0},
				{5152, 5151},
			},
			want: 2,
		},
		{
			name: "single element",
			points: [][]int{
				{1, 1},
			},
			want: 1,
		},
		{
			name: "a is not integer",
			points: [][]int{
				{1, 1},
				{3, 2},
			},
			want: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, maxPoints1(tc.points))
			assert.Equal(t, tc.want, maxPoints2(tc.points))
		})
	}
}
