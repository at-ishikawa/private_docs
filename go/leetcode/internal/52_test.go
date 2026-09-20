package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalNQueen(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "n=1",
			n:    1,
			want: 1,
		},
		{
			name: "n=4",
			n:    4,
			want: 2,
		},
		{
			name: "n=2",
			n:    2,
			want: 0,
		},
		{
			name: "n=3",
			n:    3,
			want: 0,
		},
		{
			name: "n=9",
			n:    9,
			want: 352,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := totalNQueens(tc.n)
			assert.Equal(t, tc.want, got)
		})
	}
}
