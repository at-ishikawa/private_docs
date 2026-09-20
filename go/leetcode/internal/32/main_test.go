package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestValidParentheses(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example 1",
			s:    "(()",
			want: 2,
		},
		{
			name: "example 2",
			s:    ")()())",
			want: 4,
		},
		{
			name: "example 3",
		},
		{
			name: "custom",
			s:    "((()))()",
			want: 8,
		},
		{
			name: "test case 1",
			s:    "()(()",
			want: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, longestValidParentheses1(tc.s))
			assert.Equal(t, tc.want, longestValidParentheses2(tc.s))
		})
	}
}
