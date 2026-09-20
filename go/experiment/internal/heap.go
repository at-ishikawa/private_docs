package internal

type Node struct {
	Value int

	Left      *Node
	LeftCount int

	Right      *Node
	RightCount int
}

type MinHeap Node

func (h MinHeap) Push(value int) {
	var tmp int
	if h.Value < value {
		// new node is a root value
		tmp = h.Value
		h.Value = value
	} else {
		tmp = value
	}

	if h.Left == nil {
		h.Left = &Node{
			Value: tmp,
		}
		h.LeftCount++
		return
	}
	var node *Node
	if tmp < h.Left.Value {
		// insert to the left
		node = h.Left
		h.LeftCount++
	} else if h.Right == nil {
		h.Right = &Node{
			Value: tmp,
		}
		h.RightCount++
		return
	} else {
		node = h.Right
		h.RightCount++
	}

	for node != nil {
		if node.Value < tmp {
			if node.Right == nil {
				node.Right = &Node{
					Value: tmp,
				}
				break
			}
			node = node.Right
		} else {
			if node.Left == nil {
				node.Left = &Node{
					Value: tmp,
				}
				break
			}
			node = node.Left
		}
	}
}

func MinValue(node *Node) int {
	if node.Left == nil {
		return node.Value
	}
	return MinValue(node.Left)
}

func (h *MinHeap) Pop() int {
	result := h.Value
	if h.Left != nil {
		h.Value = MinValue((*Node)(h))
		h.LeftCount--
	} else if h.RightCount > 0 {
		h.Value = MinValue(h.Right)
		h.RightCount--
	}
	return result
}

func main() {

}
