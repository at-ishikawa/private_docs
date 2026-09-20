package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	testCases := map[string]struct {
		s    string
		want int
	}{
		"Example 1": {
			s:    "1 + 1",
			want: 2,
		},
		"Example 2": {
			s:    " 2-1 + 2 ",
			want: 3,
		},
		"Example 3": {
			s:    "(1+(4+5+2)-3)+(6+8)",
			want: 23,
		},

		"Example 4": {
			s: "0",
		},
		"Example 5": {
			s:    "(1)",
			want: 1,
		},
		"negative parenthesis": {
			s:    "(1-(4+5+2)-3)-(6+8)",
			want: -27,
		},
		"simple multiply": {
			s:    "2 * 4",
			want: 8,
		},
		"multiply": {
			s:    "-(2*3)*(3+4)",
			want: -42,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := calculate(tc.s)
			assert.Equal(t, tc.want, got)
		})
	}
}
