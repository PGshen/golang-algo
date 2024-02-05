/*
 * @Date: 2024-02-05 23:32:34
 * @LastEditors: PGshen pgs1108pgs@126.com
 * @LastEditTime: 2024-02-05 23:34:51
 * @FilePath: /golang-algo/sort/insertion_sort.go
 */
package sort

// 插入排序
func insertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		base := nums[i]
		j := i - 1
		// 比base小，则往后移动
		for j >= 0 && nums[j] > base {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = base
	}
}
