package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxPathSum(t *testing.T) {
	testCases := map[string]struct {
		root *TreeNode
		want int
	}{
		"example 1": {
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: 6,
		},
		"example 2": {
			root: &TreeNode{
				Val: -10,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			want: 42,
		},

		"Pick up the sum of root nodes": {
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: -20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			want: 19,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := maxPathSum(tc.root)
			assert.Equal(t, tc.want, got)
		})
	}
}
