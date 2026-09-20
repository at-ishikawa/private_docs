package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMovesToCaptureTheQueen(t *testing.T) {
	testCases := []struct {
		name string
		a    int
		b    int
		c    int
		d    int
		e    int
		f    int
		want int
	}{
		{
			name: "example 1",

			a: 1, b: 1,
			c: 8, d: 8,
			e: 2, f: 3,

			want: 2,
		},
		{
			name: "example 2",

			a: 5, b: 3,
			c: 3, d: 4,
			e: 5, f: 2,

			want: 1,
		},
		{
			name: "test case 1",

			a: 6, b: 8,
			c: 6, d: 6,
			e: 6, f: 3,

			want: 2,
		},
		{
			name: "test case 2",

			a: 4, b: 3,
			c: 3, d: 4,
			e: 5, f: 2,

			want: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, minMovesToCaptureTheQueen1(tc.a, tc.b, tc.c, tc.d, tc.e, tc.f))
			assert.Equal(t, tc.want, minMovesToCaptureTheQueen2(tc.a, tc.b, tc.c, tc.d, tc.e, tc.f))
		})
	}
}
