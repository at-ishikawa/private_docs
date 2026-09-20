package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWords(t *testing.T) {
	testCases := []struct {
		name  string
		board [][]byte
		words []string
		want  []string
	}{
		{
			name: "example 1",
			board: [][]byte{
				[]byte("oaan"),
				[]byte("etae"),
				[]byte("ihkr"),
				[]byte("iflv"),
			},
			words: []string{"oath", "pea", "eat", "rain"},
			want: []string{
				"eat",
				"oath",
			},
		},
		{
			name: "example 2",
			board: [][]byte{
				[]byte("ab"),
				[]byte("cd"),
			},
			words: []string{
				"abcb",
			},
		},
		{
			name: "the same letter in a different cell",
			board: [][]byte{
				[]byte("oaan"),
				[]byte("etae"),
				[]byte("ihkr"),
				[]byte("iflv"),
				[]byte("deer"),
			},
			words: []string{"oath", "pea", "eat", "rain", "deer"},
			want: []string{
				"eat",
				"oath",
				"deer",
			},
		},
		{
			name: "single case",
			board: [][]byte{
				[]byte("a"),
			},
			words: []string{"a"},
			want: []string{
				"a",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := findWords(tc.board, tc.words)
			assert.ElementsMatch(t, tc.want, got)
		})
	}
}
