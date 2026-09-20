package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfAlternatingGroups(t *testing.T) {
	testCases := []struct {
		name    string
		colors  []int
		queries [][]int
		want    []int
	}{
		{
			name: "example 1",
			colors: []int{
				0, 1, 1, 0, 1,
			},
			queries: [][]int{
				{2, 1, 0},
				{1, 4},
			},
			want: []int{2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, numberOfAlternatingGroups(tc.colors, tc.queries))
		})
	}
}
