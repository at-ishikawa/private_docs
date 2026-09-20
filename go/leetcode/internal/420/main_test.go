package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrongPasswordChecker(t *testing.T) {
	testCases := []struct {
		name     string
		password string
		want     int
	}{
		{
			name:     "example 1",
			password: "a",
			want:     5,
		},
		{
			name:     "example 2",
			password: "aA1",
			want:     3,
		},
		{
			name:     "example 3",
			password: "1337C0d3",
			want:     0,
		},
		{
			name:     "test case 1",
			password: "aaa123",
			want:     1,
		},
		{
			name:     "test case 3",
			password: "aaa111",
			want:     2,
		},
		{
			name:     "test case 3",
			password: "1111111111",
			want:     3,
		},
		{
			name:     "test case 4",
			password: "ABABABABABABABABABAB1",
			want:     2,
		},
		{
			name:     "test case 5",
			password: "bbaaaaaaaaaaaaaaacccccc",
			want:     8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, strongPasswordChecker(tc.password))
		})
	}
}
