package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBattleships(t *testing.T) {
	testCases := []struct {
		name  string
		board [][]byte
		want  int
	}{
		{
			name: "example 1",
			board: [][]byte{
				{'X', '.', '.', 'X'},
				{'.', '.', '.', 'X'},
				{'.', '.', '.', 'X'},
			},
			want: 2,
		},
		{
			name: "example 2",
			board: [][]byte{
				{'.'},
			},
			want: 0,
		},
		{
			name: "example 3",
			board: [][]byte{
				{'X', '.', '.', '.'},
				{'.', '.', '.', 'X'},
				{'.', '.', '.', 'X'},
			},
			want: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, countBattleships(tc.board))
		})
	}
}
