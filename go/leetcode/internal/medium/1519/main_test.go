package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSubTrees(t *testing.T) {
	testCases := []struct {
		name   string
		n      int
		edges  [][]int
		labels string
		want   []int
	}{
		{
			name: "example 1",
			n:    7,
			edges: [][]int{
				{0, 1},
				{0, 2},
				{1, 4},
				{1, 5},
				{2, 3},
				{2, 6},
			},
			labels: "abaedcd",
			want:   []int{2, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "example 2",
			n:    4,
			edges: [][]int{
				{0, 1},
				{1, 2},
				{0, 3},
			},
			labels: "bbbb",
			want:   []int{4, 2, 1, 1},
		},
		{
			name: "example 3",
			n:    5,
			edges: [][]int{
				{0, 1},
				{0, 2},
				{1, 3},
				{0, 4},
			},
			labels: "aabab",
			want:   []int{3, 2, 1, 1, 1},
		},
		{
			name: "test case 1",
			n:    4,
			// 0
			// - 2
			//   - 1
			// - 3
			edges: [][]int{
				{0, 2},
				{0, 3},
				{1, 2},
			},
			labels: "aeed",
			want:   []int{1, 1, 2, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, countSubTrees1(tc.n, tc.edges, tc.labels))
			assert.Equal(t, tc.want, countSubTrees2(tc.n, tc.edges, tc.labels))

		})
	}
}
