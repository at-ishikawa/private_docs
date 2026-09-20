// https://leetcode.com/problems/construct-string-from-binary-tree/description/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS. Generate parenthesis and return to parents
// Cases for each node
// 1. No child: Node
// 2. No left child: Node(()(Right child))
// 3. No right child: Node(Left Child)
// 4. Both children: Node(Left child)(Right child)

// Time complexity: O(N)
// Space complexity: O(1)
func tree2str(root *TreeNode) string {
	builder := strings.Builder{}
	builder.WriteString(strconv.Itoa(root.Val))

	if root.Left != nil {
		builder.WriteString(fmt.Sprintf("(%s)", tree2str(root.Left)))
	}
	if root.Right != nil {
		if root.Left == nil {
			builder.WriteString("()")
		}
		builder.WriteString(fmt.Sprintf("(%s)", tree2str(root.Right)))
	}
	return builder.String()
}
