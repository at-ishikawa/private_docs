package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxScore(t *testing.T) {
	testCases := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "example 1",
			grid: [][]int{
				{1, 2, 3},
				{4, 3, 2},
				{1, 1, 1},
			},
			want: 8,
		},
		{
			name: "example 2",
			grid: [][]int{
				{8, 7, 6},
				{8, 3, 2},
			},
			want: 15,
		},
		{
			name: "my test case",
			grid: [][]int{
				{1},
				{1},
			},
			want: 1,
		},
		{
			name: "test case 1",
			grid: [][]int{
				{16, 18},
				{20, 20},
				{18, 18},
				{1, 15},
			}, // 20, 18, 16, 15,
			want: 69,
		},
		{
			name: "test case 2",
			grid: [][]int{
				{16, 14},
				{5, 4},
				{7, 16},
			},
			want: 35,
		},
		{
			name: "time limit exceeded",
			grid: [][]int{
				{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: 55,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// assert.Equal(t, tc.want, maxScore1(tc.grid))
			assert.Equal(t, tc.want, maxScore2(tc.grid))
		})
	}
}
