package sort

// 选择排序
func selectionSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		k := i
		// 选取最小的跟i位置交换
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[k] {
				k = j
			}
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
}
