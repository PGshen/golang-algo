package sort

// 希尔排序，插入排序优化版
func shellSort(nums []int) {
	n := len(nums)
	// 增量序列：初始增量设为数组的一半，之后每次减半
	for gap := n / 2; gap > 0; gap /= 2 {
		// 对每个子序列进行插入排序
		for i := gap; i < n; i++ {
			for j := i; j >= gap && nums[j-gap] > nums[j]; j -= gap {
				nums[j], nums[j-gap] = nums[j-gap], nums[j]
			}
		}
	}
}
