package doublepointer

import (
	"fmt"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	node3 := ListNode{
		Val:  3,
		Next: nil,
	}
	node2 := ListNode{
		Val:  2,
		Next: nil,
	}
	node0 := ListNode{
		Val:  0,
		Next: nil,
	}
	node4 := ListNode{
		Val:  -4,
		Next: nil,
	}
	node3.Next = &node2
	node2.Next = &node0
	node0.Next = &node4
	node4.Next = &node2
	fmt.Println(detectCycle(&node3).Val)
}
