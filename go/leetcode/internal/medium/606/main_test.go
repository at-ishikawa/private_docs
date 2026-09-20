package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree2Str(t *testing.T) {
	testCases := []struct {
		name string
		root *TreeNode
		want string
	}{
		{
			name: "example 1",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: "1(2(4))(3)",
		},
		{
			name: "example 2",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: "1(2()(4))(3)",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tree2str(tc.root))
		})
	}
}
