package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyRandomList(t *testing.T) {
	testCases := []struct {
		name     string
		nodeFunc func() *Node
	}{
		{
			name: "example 3",
			nodeFunc: func() *Node {
				node := &Node{
					Val: 3,
					Next: &Node{
						Val: 3,
						Next: &Node{
							Val: 3,
						},
					},
				}
				node.Next.Random = node
				return node
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Run("copyRandomList1", func(t *testing.T) {
				arg := tc.nodeFunc()
				got := copyRandomList1(arg)
				assert.NotSame(t, arg, got)
				assert.Equal(t, arg, got)
			})
			t.Run("copyRandomList2", func(t *testing.T) {
				arg := tc.nodeFunc()
				got := copyRandomList2(arg)
				assert.NotSame(t, arg, got)
				assert.Equal(t, arg, got)
			})
		})
	}
}
