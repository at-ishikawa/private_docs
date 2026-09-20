package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMergeLists(t *testing.T) {
	testCases := []struct {
		name  string
		lists []*ListNode
		want  *ListNode
	}{
		{
			name: "[1, 4, 5],[1,3,4],[2,6]",
			lists: []*ListNode{
				{
					Val: 1,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{Val: 5},
					},
				},
				{
					Val: 1,
					Next: &ListNode{
						Val:  3,
						Next: &ListNode{Val: 4},
					},
				},
				{
					Val: 2,
					Next: &ListNode{
						Val: 6,
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val:  5,
										Next: &ListNode{Val: 6},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "empty",
			lists: nil,
			want:  nil,
		},
		{
			name:  "empty list",
			lists: nil,
			want:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			backup := func(lists []*ListNode) []*ListNode {
				if len(lists) == 0 {
					return lists
				}

				result := make([]*ListNode, 0)
				for _, node := range lists {
					current := &ListNode{}
					result = append(result, current)
					var previousNode *ListNode
					for ; node != nil; node = node.Next {
						current.Val = node.Val
						current.Next = &ListNode{}
						previousNode = current
						current = current.Next
					}
					if previousNode != nil {
						previousNode.Next = nil
					}
				}
				return result
			}

			t.Run("mergeKLists1", func(t *testing.T) {
				lists := backup(tc.lists)
				require.Equal(t, tc.lists, lists)
				got := mergeKLists1(lists)
				assert.Equal(t, tc.want, got)
			})

			t.Run("mergeKLists2", func(t *testing.T) {
				lists := backup(tc.lists)
				require.Equal(t, tc.lists, lists)
				got := mergeKLists2(lists)
				assert.Equal(t, tc.want, got)
			})

			t.Run("mergeKLists3", func(t *testing.T) {
				lists := backup(tc.lists)
				require.Equal(t, tc.lists, lists)
				got := mergeKLists3(lists)
				assert.Equal(t, tc.want, got)
			})
		})
	}
}
