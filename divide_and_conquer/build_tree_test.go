package divideandconquer

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 2, 1, 7}
	inorder := []int{9, 3, 1, 2, 7}
	tree := buildTree(preorder, inorder)
	fmt.Printf("%v", tree)
}

func TestBuildTree2(t *testing.T) {
	postorder := []int{9, 1, 7, 2, 3}
	inorder := []int{9, 3, 1, 2, 7}
	tree := buildTree2(postorder, inorder)
	fmt.Printf("%v", tree)
}
