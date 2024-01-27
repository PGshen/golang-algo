package datastruct

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	// root := newBSearchTreeNode(8)
	var root *bSearchTree = newBSearchTree()
	root.insert(8)
	root.insert(4)
	root.insert(12)
	root.insert(2)
	root.insert(6)
	root.insert(10)
	root.insert(14)
	fmt.Printf("%v", root.inOrder()...)
	fmt.Println()
	root.remove(8)
	fmt.Printf("%v", root.inOrder()...)
}
