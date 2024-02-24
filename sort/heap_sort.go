package sort

// 堆排序
func heapSort(nums []int) {
	// 建堆
	for i := len(nums)/2 - 1; i >= 0; i-- {
		siftDown(&nums, len(nums), i)
	}
	// 从堆中提取最大值，循环n-1轮
	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		siftDown(&nums, i, 0)
	}
}

// 向下堆化
func siftDown(nums *[]int, n, i int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		maxI := i
		// 交换左右子节点的大值
		if left < n && (*nums)[left] > (*nums)[maxI] {
			maxI = left
		}
		if right < n && (*nums)[right] > (*nums)[maxI] {
			maxI = right
		}
		if maxI == i {
			break
		}
		(*nums)[i], (*nums)[maxI] = (*nums)[maxI], (*nums)[i]
		i = maxI
	}
}
