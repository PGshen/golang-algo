package search

import (
	"container/heap"

	"github.com/PGshen/golang-algo/datastruct"
)

func findKthLargest(nums []int, k int) int {
	minHeap := &datastruct.MinHeap{}
	heap.Init(minHeap)
	i := 0
	for i < k {
		heap.Push(minHeap, nums[i])
		i++
	}
	for i < len(nums) {
		heap.Push(minHeap, nums[i])
		heap.Pop(minHeap)
		i++
	}
	return minHeap.Top().(int)
}

// 快排算法变种
func findKthLargest2(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for {
		if left >= right {
			return nums[right]
		}
		p := partition(nums, left, right)
		if p+1 == k {
			return nums[p]
		} else if p+1 < k {
			left = p + 1
		} else {
			right = p - 1
		}
	}
}

func partition(nums []int, left int, right int) int {
	pivot := nums[right]
	for i := left; i < right; i++ {
		if nums[i] > pivot {
			nums[left], nums[i] = nums[i], nums[right]
			left++
		}
	}
	nums[left], nums[right] = nums[right], nums[left]
	return left
}
