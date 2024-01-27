package datastruct

import (
	"fmt"
	"testing"
)

func TestArrayBinaryTree(t *testing.T) {
	arr := []any{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	bt := newArrayBinaryTree(arr)
	fmt.Printf("%v", bt.levelOrder()...)
	fmt.Println()

	fmt.Printf("%v", bt.preOrder()...)
	fmt.Println()

	fmt.Printf("%v", bt.inOrder()...)
	fmt.Println()

	fmt.Printf("%v", bt.postOrder()...)
}
