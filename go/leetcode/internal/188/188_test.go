package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		name   string
		k      int
		prices []int
		want   int
	}{
		{
			name:   "example 1",
			k:      2,
			prices: []int{2, 4, 1},
			want:   2,
		},
		{
			name:   "example 2",
			k:      2,
			prices: []int{3, 2, 6, 5, 0, 3},
			want:   7,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, maxProfit1(tc.k, tc.prices))
			assert.Equal(t, tc.want, maxProfit2(tc.k, tc.prices))
		})
	}
}
