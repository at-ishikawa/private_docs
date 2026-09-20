package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinWindow(t *testing.T) {
	testCases := map[string]struct {
		s    string
		t    string
		want string
	}{
		"example 1": {
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		"example 2": {
			s:    "a",
			t:    "a",
			want: "a",
		},
		"example 3": {
			s: "a",
			t: "aa",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := minWindow_76(tc.s, tc.t)
			assert.Equal(t, tc.want, got)
		})
	}
}
