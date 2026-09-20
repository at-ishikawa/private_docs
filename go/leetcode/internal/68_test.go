package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullJustify(t *testing.T) {
	testCases := []struct {
		name     string
		words    []string
		maxWidth int
		want     []string
	}{
		{
			name: "example 1",
			words: []string{
				"This", "is", "an", "example", "of", "text", "justification.",
			},
			maxWidth: 16,
			want: []string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		},
		{
			name: "example 2",
			words: []string{
				"What", "must", "be", "acknowledgment", "shall", "be",
			},
			maxWidth: 16,
			want: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		},
		{
			name: "example 3",
			words: []string{
				"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do",
			},
			maxWidth: 20,
			want: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, fullJustify1(tc.words, tc.maxWidth))
		})
	}
}
