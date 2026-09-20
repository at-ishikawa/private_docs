package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDigitOne(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "example 1",
			n:    13,
			want: 6,
		},
		{
			name: "example 2",
			n:    0,
			want: 0,
		},
		{
			name: "test case 1",
			n:    100,
			want: 21,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, countDigitOne1(tc.n))
			assert.Equal(t, tc.want, countDigitOne2(tc.n))
		})
	}
}
