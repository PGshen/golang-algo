package datastruct

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	maxHeap := &intHeap{}
	heap.Init(maxHeap)

	heap.Push(maxHeap, 1)
	heap.Push(maxHeap, 3)
	heap.Push(maxHeap, 2)
	heap.Push(maxHeap, 4)
	heap.Push(maxHeap, 5)
	heap.Push(maxHeap, 8)
	heap.Push(maxHeap, 6)

	top := maxHeap.Top()
	fmt.Println(top)

	heap.Pop(maxHeap)
	fmt.Println(maxHeap.Top())

}
