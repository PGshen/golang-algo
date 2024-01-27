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
	fmt.Printf("%v", bt.LevelOrder()...)
	fmt.Println()

	fmt.Printf("%v", bt.PreOrder()...)
	fmt.Println()

	fmt.Printf("%v", bt.InOrder()...)
	fmt.Println()

	fmt.Printf("%v", bt.PostOrder()...)

	bt.parent(3)
}
