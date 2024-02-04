/*
 * @Date: 2024-02-04 23:19:07
 * @LastEditors: PGshen pgs1108pgs@126.com
 * @LastEditTime: 2024-02-04 23:26:04
 * @FilePath: /golang-algo/sort/bubble_sort.go
 */
package sort

// 冒泡排序
func bubbleSort(nums []int) {
	n := len(nums)
	// 未排序区 [0, i]
	for i := n - 1; i > 0; i-- {
		// 未排序区 [0, i] 的最大元素交换到最右端
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// 冒泡排序，优化版
func bubbleSortWithFlag(nums []int) {
	n := len(nums)
	// 未排序区 [0, i]
	for i := n - 1; i > 0; i-- {
		flag := false
		// 未排序区 [0, i] 的最大元素交换到最右端
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		// 此轮“冒泡”未交换任何元素, 说明已经有序了，提前退出
		if !flag {
			break
		}
	}
}
