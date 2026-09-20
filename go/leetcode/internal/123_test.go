package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name: "example 1",
			prices: []int{
				3, 3, 5, 0, 0, 3, 1, 4,
			},
			want: 6,
		},
		{
			name: "example 2",
			prices: []int{
				1, 2, 3, 4, 5,
			},
			want: 4,
		},
		{
			name: "example 3",
			prices: []int{
				7, 6, 4, 3, 1,
			},
			want: 0,
		},
		{
			name: "no triple transactions",
			prices: []int{
				1, 3, 1, 4, 1, 5,
			},
			want: 7,
		},
		{
			name:   "one day",
			prices: []int{1},
			want:   0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, maxProfit(tc.prices))
		})
	}
}
