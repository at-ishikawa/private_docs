package internal

import (
	"container/heap"
	"math"
)

func mergeKLists(lists []*ListNode) *ListNode {
	return mergeKLists2(lists)
}

func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	head := &ListNode{}
	current := head
	var previousNode *ListNode
	for true {
		minVal := math.MaxInt
		minIndex := 0

		for index, list := range lists {
			if list == nil {
				continue
			}
			firstElement := list.Val
			if firstElement < minVal {
				minIndex = index
				minVal = firstElement
			}
		}
		if minVal == math.MaxInt {
			break
		}

		lists[minIndex] = lists[minIndex].Next
		current.Val = minVal
		current.Next = &ListNode{}
		previousNode = current
		current = current.Next
	}

	if previousNode == nil {
		return nil
	}

	previousNode.Next = nil
	return head
}

type KListsPriorityQueue struct {
	nodes []*ListNode
}

var _ heap.Interface = (*KListsPriorityQueue)(nil)

// Len implements heap.Interface.
func (k *KListsPriorityQueue) Len() int {
	return len(k.nodes)
}

// Less implements heap.Interface.
func (k *KListsPriorityQueue) Less(i int, j int) bool {
	// fmt.Printf("%d, %d\n", i, j)
	return k.nodes[i].Val < k.nodes[j].Val
}

// Pop implements heap.Interface.
func (k *KListsPriorityQueue) Pop() any {
	node := k.nodes[0]
	k.nodes = k.nodes[1:]
	return node
}

// Push implements heap.Interface.
func (k *KListsPriorityQueue) Push(x any) {
	k.nodes = append(k.nodes, x.(*ListNode))
}

// Swap implements heap.Interface.
func (k *KListsPriorityQueue) Swap(i int, j int) {
	k.nodes[i], k.nodes[j] = k.nodes[j], k.nodes[i]
}

// Use a priority queue
// Space: O(k)
func mergeKLists2(lists []*ListNode) *ListNode {
	queue := &KListsPriorityQueue{
		nodes: make([]*ListNode, 0),
	}
	for _, headNode := range lists {
		queue.nodes = append(queue.nodes, headNode)
		// for ; list != nil; list = list.Next {
		//			queue.nodes = append(queue.nodes, list)
		//}
	}
	heap.Init(queue)
	if queue.Len() == 0 {
		return nil
	}

	head := &ListNode{}
	current := head
	for queue.Len() > 0 {
		node := heap.Pop(queue).(*ListNode)
		current.Next = node
		current = node
		if node.Next != nil {
			heap.Push(queue, node.Next)
		}
	}

	return head.Next
}

// Use a divide and conquer
func mergeKLists3(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	mid := len(lists)
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])
	head := &ListNode{}
	node := head
	for left != nil && right != nil {
		if left.Val < right.Val {
			node.Next = left
			node = left
		} else {
			node.Next = right
			node = right
		}
	}
	if left == nil {
		node.Next = right
	} else {
		node.Next = left
	}

	return head.Next
}
