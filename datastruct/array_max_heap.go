package datastruct

import "fmt"

type arrayMaxHeap struct {
	data []any
}

func newArrayMaxHeap() *arrayMaxHeap {
	return &arrayMaxHeap{
		data: make([]any, 0),
	}
}

// 根据切片建堆
func newMaxHeap(nums []any) *arrayMaxHeap {
	h := &arrayMaxHeap{
		data: nums,
	}
	for i := h.parent(h.size() - 1); i >= 0; i-- {
		h.siftDown(i)
	}
	return h
}

func (h *arrayMaxHeap) left(i int) int {
	return 2*i + 1
}

func (h *arrayMaxHeap) right(i int) int {
	return 2*i + 2
}

func (h *arrayMaxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *arrayMaxHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *arrayMaxHeap) size() int {
	return len(h.data)
}

func (h *arrayMaxHeap) isEmpty() bool {
	return h.size() == 0
}

func (h *arrayMaxHeap) peek() any {
	return h.data[0]
}

func (h *arrayMaxHeap) push(value any) {
	// 放入最后
	h.data = append(h.data, value)
	// 向上堆化
	h.siftUp(h.size() - 1)
}

func (h *arrayMaxHeap) siftUp(i int) {
	for {
		p := h.parent(i)
		if p < 0 || h.data[i].(int) <= h.data[p].(int) {
			break
		}
		h.swap(i, p)
		// 循环向上堆化
		i = p
	}
}

func (h *arrayMaxHeap) pop() any {
	if h.isEmpty() {
		fmt.Println("empty")
		return nil
	}
	// 根节点和最后节点交换
	h.swap(0, h.size()-1)
	val := h.data[h.size()-1]
	// 删除最后节点
	h.data = h.data[:h.size()-1]
	// 根节点向下堆化
	h.siftDown(0)
	return val
}

func (h *arrayMaxHeap) siftDown(i int) {
	for {
		l, r, max := h.left(i), h.right(i), i
		if l < h.size() && h.data[l].(int) > h.data[max].(int) {
			max = l
		}
		if r < h.size() && h.data[r].(int) > h.data[max].(int) {
			max = r
		}
		// 节点i已经最大
		if max == i {
			break
		}
		h.swap(i, max)
		// 循环向下堆化
		i = max
	}
}
