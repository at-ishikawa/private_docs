package internal

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time complexity: O(N + N) => O(N)
// Space complexity: O(1)

func reverseKGroup(head *ListNode, k int) *ListNode {
	nextHead := head
	var i int
	for ; i < k && nextHead != nil; i++ {
		nextHead = nextHead.Next
	}
	if i != k {
		return head
	}
	nextHead = reverseKGroup(nextHead, k)

	current := head
	var prevNode *ListNode
	for i := 0; i < k; i++ {
		nextNode := current.Next
		current.Next = prevNode
		prevNode = current
		current = nextNode
	}
	head.Next = nextHead
	return prevNode
	/*
		// Non recursive solution
			var prevNode *ListNode
			var prevHead *ListNode
			current := head
			var newHead *ListNode
			for current != nil {
				kTail := current
				i := 0
				for ; i < k && kTail != nil; i++ {
					kTail = kTail.Next
				}
				if i != k {
					if prevHead != nil {
						prevHead.Next = current
					}
					break
				}

				// there are k elements from current to kTail
				// b -> a, c -> d -> e
				// current: c
				// prevNode: b
				// head: a
				// b -> a -> d -> c, e
				// current: e
				// prevNode: d
				//
				kHead := current
				prevNode = nil
				for current != kTail {
					nextNode := current.Next
					current.Next = prevNode
					prevNode = current
					current = nextNode
				}
				if newHead == nil {
					newHead = prevNode
				} else {
					prevHead.Next = prevNode
				}
				prevHead = kHead
				// fmt.Printf("Current: %+v, prev node: %+v, prevHead: %+v\n", current, prevNode, prevHead)
			}
			if newHead == nil {
				return head
			}
			return newHead
	*/
}
