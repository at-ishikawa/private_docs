package internal

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	return copyRandomList2(head)
}

func copyRandomList2(head *Node) *Node {
	oldToNewNodeMap := make(map[*Node]*Node)
	for node := head; node != nil; node = node.Next {
		oldToNewNodeMap[node] = &Node{
			Val: node.Val,
		}
	}

	for node := head; node != nil; node = node.Next {
		newNode := oldToNewNodeMap[node]
		newNode.Next = oldToNewNodeMap[node.Next]
		newNode.Random = oldToNewNodeMap[node.Random]
	}
	return oldToNewNodeMap[head]
}

func copyRandomList1(head *Node) *Node {
	nodeIndexMap := make(map[*Node]int, 0) // pointer => position
	index := 0
	for node := head; node != nil; node = node.Next {
		nodeIndexMap[node] = index
		index++
	}

	// the nth node => ith node
	randoms := make([]int, index)
	newNodes := make([]*Node, index)

	index = 0
	for node := head; node != nil; node = node.Next {
		randoms[index] = -1
		if node.Random != nil {
			destNodeIndex, ok := nodeIndexMap[node.Random]
			if ok {
				randoms[index] = destNodeIndex
			}
		}

		newNodes[index] = &Node{
			Val: node.Val,
		}
		index++
	}

	newHead := &Node{}
	current := newHead
	for i := 0; i < index; i++ {
		nextNode := newNodes[i]
		current.Next = nextNode
		if randoms[i] >= 0 {
			nextNode.Random = newNodes[randoms[i]]
		}
		current = nextNode
	}

	return newHead.Next
}
