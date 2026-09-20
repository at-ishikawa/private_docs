package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxIceCream(t *testing.T) {
	testCases := []struct {
		name  string
		costs []int
		coins int
		want  int
	}{
		{
			name: "example 1",
			costs: []int{
				1, 3, 2, 4, 1,
			},
			// 1 3 2 4 1
			// 0 2 1 1 1
			// 0 2 3 4 5
			coins: 7,
			want:  4,
		},
		{
			name: "example 2",
			costs: []int{
				10, 6, 8, 7, 7, 8,
			},
			coins: 5,
			want:  0,
		},
		{
			name: "example 3",
			costs: []int{
				1, 6, 3, 1, 2, 5,
			},
			coins: 20,
			want:  6,
		},
		{
			name: "test case 1",
			costs: []int{
				7, 3, 3, 6, 6, 6, 10, 5, 9, 2,
				// 0, 0, 1, 2, 0, 1, 3, 1, 0, 1
				// 0, 0, 1, 3, 3, 4, 7, 8, 8, 9
			},
			coins: 56,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, maxIceCream(tc.costs, tc.coins))
		})
	}
}
