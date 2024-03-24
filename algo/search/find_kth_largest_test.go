package search

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest2([]int{3, 2, 1, 5, 6, 4}, 2))
}
