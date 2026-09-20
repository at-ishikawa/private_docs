package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxmizedCapital(t *testing.T) {
	type args struct {
		k       int
		w       int
		profits []int
		capital []int
	}
	testCases := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				k:       2,
				w:       0,
				profits: []int{1, 2, 3},
				capital: []int{0, 1, 1},
			},
			want: 4,
		},
		{
			name: "example 2",
			args: args{
				k:       3,
				w:       0,
				profits: []int{1, 2, 3},
				capital: []int{0, 1, 1},
			},
			want: 6,
		},
		{
			name: "example 3",
			args: args{
				k:       1,
				w:       2,
				profits: []int{1, 2, 3},
				capital: []int{1, 1, 2},
			},
			want: 5,
		},
		{
			name: "The minimum case",
			args: args{
				k:       1,
				w:       0,
				profits: []int{0},
				capital: []int{0},
			},
			want: 0,
		},
		{
			name: "The project select case",
			args: args{
				k:       2,
				w:       1,
				profits: []int{3, 1, 2},
				capital: []int{3, 1, 1},
			},
			want: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := findMaximizedCapital(tc.args.k, tc.args.w, tc.args.profits, tc.args.capital)
			assert.Equal(t, tc.want, got)
		})
	}
}
