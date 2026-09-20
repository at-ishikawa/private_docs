// 1261. Find Elements in a Contaminated Binary Tree:
// https://leetcode.com/problems/find-elements-in-a-contaminated-binary-tree/description/
package main

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	root *TreeNode
}

func Constructor(root *TreeNode) FindElements {
	/*
		queue := []*TreeNode{
			root,
		}
		root.Val = 0
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				node.Left.Val = node.Val*2 + 1
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				node.Right.Val = node.Val*2 + 2
				queue = append(queue, node.Right)
			}
		}
	*/
	return FindElements{
		root: root,
	}
}

// the depth N: 2^(N-1) - 1
// the root: 2^(1-1)-1 => 0

// 0 (b0)
// - 1 (b0001)
//   - 3 (b0011)
//     - 7 (b0111)
//       - 15 (b1111)
//     - 8 (b1000)
//   - 4 (b0100)
//     - 9 (b1001)
//     - 10 (b1010)
//
// - 2 (b0010)
//   - 5 (b0101)
//     - 11 (b1011)
//     - 12 (b1100)
//   - 6 (b0110)
//     - 13 (b1101)
//     - 14 (b1110)

// +1
// 1(b1)
// - 2 (b10)
//   - 4 (b100)
//     - 8 (b1000)
//     - 9 (b1001)
//   - 5 (b101)
//     - 10 (b1010)
//     - 11 (b1011)
// - 3 (b11)
//   - 6 (b0110)
//     - 12 (b1100)
//     - 13 (b1101)
//   - 7 (b0111)
//     - 14 (b1110)
//     - 15 (b1111)

// 5 (=6 for 5 + 1, b0110)
// %2 => get the least big
// /2 => shift one bit to right

// 5%2 => 1 => 3, 5 / 2 => 2
// 2%2 => 0 => 6, 2 / 2 => 1
// 1/2 => 6?

// 10%2 => 0 => 2, 10 / 2 => 5
// 5 % 2 => 1 => 5, 5 / 2 => 2
// 2 % 2 => 0, 10, 2 / 1 => 1

// 13%2 => 1 => 3, 6
// 6 % 2 => 0 => 6, 6 / 2 = 3
// 3 % 2 => 1 => 13, 3 / 2 = 1

func (this *FindElements) Find(target int) bool {
	node := this.root

	binary := strconv.FormatInt(int64(target+1), 2)
	// fmt.Printf("%+v\n", binary)
	for index := 1; index < len(binary); index++ {
		currentBit := int(binary[index]) & 1
		// fmt.Printf("%d, %d, %v\n", index, int(binary[index]), currentBit)
		if currentBit == 1 {
			node = node.Right
		} else {
			node = node.Left
		}

		if node == nil {
			return false
		}
	}

	return true
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
