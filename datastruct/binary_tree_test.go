package datastruct

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	bRoot := newBTreeNode("root")
	n1 := newBTreeNode("1")
	n2 := newBTreeNode("2")
	n3 := newBTreeNode("3")
	n4 := newBTreeNode("4")
	n5 := newBTreeNode("5")
	n6 := newBTreeNode("6")

	bRoot.Left = n1
	bRoot.Right = n2
	n1.Left = n3
	n1.Right = n4
	n2.Left = n5
	n2.Right = n6

	fmt.Printf("%v", bRoot.levelOrder()...)
	fmt.Println()

	fmt.Printf("%v", bRoot.preOrder([]any{})...)
	fmt.Println()

	fmt.Printf("%v", bRoot.inOrder([]any{})...)
	fmt.Println()

	fmt.Printf("%v", bRoot.postOrder([]any{})...)
}
