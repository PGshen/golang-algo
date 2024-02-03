package datastruct

import (
	"fmt"
	"testing"
)

func TestArrayMaxHeap(t *testing.T) {
	maxHeap := newArrayMaxHeap()
	maxHeap.push(1)
	maxHeap.push(3)
	maxHeap.push(6)
	maxHeap.push(2)
	maxHeap.push(9)
	maxHeap.push(2)
	maxHeap.push(5)
	maxHeap.push(9)

	fmt.Println(maxHeap.peek())

	// maxHeap.pop()
	fmt.Println(maxHeap.pop())
	fmt.Println(maxHeap.pop())
	fmt.Println(maxHeap.pop())
	fmt.Println(maxHeap.pop())

	maxHeap2 := newMaxHeap([]any{1, 2, 5, 3, 7, 4})
	fmt.Println(maxHeap2.peek())
}
