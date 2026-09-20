package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindElements(t *testing.T) {
	type functionCall struct {
		target int
		want   bool
	}
	testCases := []struct {
		name          string
		root          *TreeNode
		functionCalls []functionCall
	}{
		{
			name: "example 1",
			root: &TreeNode{
				Val: -1,
				Right: &TreeNode{
					Val: -1,
				},
			},
			functionCalls: []functionCall{
				{target: 1},
				{target: 2, want: true},
			},
		},
		{
			name: "example 2",
			root: &TreeNode{
				Val: -1,
				Left: &TreeNode{
					Val: -1,
					Left: &TreeNode{
						Val: -1,
					},
					Right: &TreeNode{
						Val: -1,
					},
				},
				Right: &TreeNode{
					Val: -1,
				},
			},
			functionCalls: []functionCall{
				{target: 1, want: true},
				{target: 3, want: true},
				{target: 5},
			},
		},
		{
			name: "example 3",
			root: &TreeNode{
				Val: -1,
				Right: &TreeNode{
					Val: -1,
					Left: &TreeNode{
						Val: -1,
						Left: &TreeNode{
							Val: -1,
						},
					},
				},
			},
			functionCalls: []functionCall{
				{target: 2, want: true},
				{target: 3},
				{target: 4},
				{target: 5, want: true},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			findElement := Constructor(tc.root)
			for _, call := range tc.functionCalls {
				assert.Equal(t, call.want, findElement.Find(call.target))
			}
		})
	}
}
