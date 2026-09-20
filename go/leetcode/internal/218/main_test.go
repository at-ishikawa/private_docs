package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSkyline(t *testing.T) {
	testCases := []struct {
		name      string
		buildings [][]int
		want      [][]int
	}{
		{
			name: "example 1",
			buildings: [][]int{
				{2, 9, 10},
				{3, 7, 15},
				{5, 12, 12},
				{15, 20, 10},
				{19, 24, 8},
			},
			want: [][]int{
				{2, 10},
				{3, 15},
				{7, 12},
				{12, 0},
				{15, 10},
				{20, 8},
				{24, 0},
			},
		},
		{
			name: "example 2",
			buildings: [][]int{
				{0, 2, 3},
				{2, 5, 3},
			},
			want: [][]int{
				{0, 3},
				{5, 0},
			},
		},
		{
			name: "test case 1",
			buildings: [][]int{
				{1, 2, 1},
				{1, 2, 2},
				{1, 2, 3},
			},
			want: [][]int{
				{1, 3},
				{2, 0},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, getSkyline(tc.buildings))
		})
	}
}
