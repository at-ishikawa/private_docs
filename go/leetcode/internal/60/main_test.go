package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPermutation(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		k    int
		want string
	}{
		{
			name: "example 1",
			n:    3,
			k:    3,
			want: "213",
		},
		{
			name: "example 2",
			n:    4,
			k:    9,
			want: "2314",
		},
		{
			name: "example 3",
			n:    3,
			k:    1,
			want: "123",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, getPermutation(tc.n, tc.k))
		})
	}
}
