package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBagOfTokenScore(t *testing.T) {
	testCases := []struct {
		name   string
		tokens []int
		power  int
		want   int
	}{
		{
			name:   "example 1",
			tokens: []int{100},
			power:  50,
			want:   0,
		},
		{
			name:   "example 2",
			tokens: []int{200, 100},
			power:  150,
			want:   1,
		},
		{
			name:   "example 3",
			tokens: []int{100, 200, 300, 400},
			want:   2,
			power:  200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, bagOfTokensScore(tc.tokens, tc.power))
		})
	}
}
