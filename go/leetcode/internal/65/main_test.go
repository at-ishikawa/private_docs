package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumber(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		want bool
	}{
		{s: "0", want: true},
		{s: "e"},
		{s: "."},
		{s: "2", want: true},
		{s: "0089", want: true},
		{s: "-0.1", want: true},
		{s: "+3.14", want: true},
		{s: "4.", want: true},
		{s: "-.9", want: true},
		{s: "2e10", want: true},
		{s: "-90E3", want: true},
		{s: "3e+7", want: true},
		{s: "+6e-1", want: true},
		{s: "53.5e93", want: true},
		{s: "-123.456e789", want: true},
		{s: "abc"},
		{s: "1a"},
		{s: "1e"},
		{s: "e3"},
		{s: "99e2.5"},
		{s: "--6"},
		{s: "-+3"},
		{s: "95a54e53"},
		{s: "4e+"},
	}

	for _, tc := range testCases {
		t.Run(tc.s, func(t *testing.T) {
			assert.Equal(t, tc.want, isNumber1(tc.s))
			assert.Equal(t, tc.want, isNumber2(tc.s))
		})
	}
}
