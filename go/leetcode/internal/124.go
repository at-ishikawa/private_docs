package internal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxCount int

// Time complexity: O(N)
// Space complexity: O(1)
func maxPathSum(root *TreeNode) int {
	maxCount = -1001
	findMaxPath(root)
	return maxCount
}

func findMaxPath(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := findMaxPath(node.Left)
	right := findMaxPath(node.Right)

	currentMaxPath := node.Val + left + right
	maxCount = max(maxCount, currentMaxPath)

	return max(node.Val+max(left, right, 0), 0)
}

/*
func maxPathSum(root *TreeNode) int {
	ans1, ans2 := calculateMax(root)
	return max(ans1, ans2)
}

func calculateMax(root *TreeNode) (int, int) {
	if root == nil {
		return -1001, -1001
	}

	leftConnectedMax, leftIndependentMax := calculateMax(root.Left)
	rightConnectedMax, rightIndependentMax := calculateMax(root.Right)

	sumConnected := root.Val
	maxLeaf := max(leftConnectedMax, rightConnectedMax)
	if maxLeaf > 0 {
		sumConnected += maxLeaf
	}
	if root.Val >= 0 {
		return sumConnected, max(
			sumConnected,
			root.Val+leftConnectedMax+rightConnectedMax,
			leftIndependentMax,
			rightIndependentMax,
		)
	}
	return sumConnected, max(
		sumConnected,
		root.Val+leftConnectedMax+rightConnectedMax,
		leftIndependentMax,
		rightIndependentMax,
	)
}
*/
