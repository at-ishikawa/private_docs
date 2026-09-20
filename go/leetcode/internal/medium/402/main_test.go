package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveKDigits(t *testing.T) {
	testCases := []struct {
		name string
		num  string
		k    int
		want string
	}{
		{
			name: "example 1",
			num:  "1432219",
			k:    3,
			want: "1219",
		},
		{
			name: "example 2",
			num:  "10200",
			k:    1,
			want: "200",
		},
		{
			name: "example 3",
			num:  "10",
			k:    2,
			want: "0",
		},
		{
			name: "example 4",
			num:  "9",
			k:    1,
			want: "0",
		},
		{
			name: "example 5",
			num:  "112",
			k:    1,
			want: "11",
		},
		{
			name: "example 6",
			num:  "12",
			k:    1,
			want: "1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, removeKdigits(tc.num, tc.k))
		})
	}
}
