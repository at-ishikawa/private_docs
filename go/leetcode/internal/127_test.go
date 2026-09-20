package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLadderLength(t *testing.T) {
	testCases := map[string]struct {
		beginWord string
		endWord   string
		wordList  []string
		want      int
	}{
		"example 1": {
			beginWord: "hit",
			endWord:   "cog",
			wordList: []string{
				"hot",
				"dot",
				"dog",
				"lot",
				"log",
				"cog",
			},
			want: 5,
		},
		"example 2": {
			beginWord: "hit",
			endWord:   "cog",
			wordList: []string{
				"hot",
				"dot",
				"dog",
				"lot",
				"log",
			},
			want: 0,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := ladderLength(tc.beginWord, tc.endWord, tc.wordList)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestFunc(t *testing.T) {
	testCases := []struct {
		name string
		want int
	}{
		{
			name: "happy path",
			want: 3,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
		})
	}
}
