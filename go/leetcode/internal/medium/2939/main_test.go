package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumXorProduct(t *testing.T) {
	testCases := []struct {
		name string
		a    int64
		b    int64
		n    int
		want int
	}{
		{
			name: "example 1",
			a:    12,
			b:    5,
			n:    4,
			want: 98,
		},
		{
			name: "example 2",
			a:    6,
			b:    7,
			n:    5,
			want: 930,
		},
		{
			name: "example 3",
			a:    1,
			b:    6,
			n:    3,
			want: 12,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, maximumXorProduct(tc.a, tc.b, tc.n))
		})
	}
}
