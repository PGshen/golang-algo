package sort

// 快速排序
func quickSort(nums []int, left, right int) {
	// 尾递归优化，每次仅递归较短的子数组
	for left < right {
		pivot := partition(nums, left, right)
		if pivot-left < right-pivot {
			quickSort(nums, left, pivot-1) // 递归排序左子数组
			left = pivot + 1               // 剩余未排序区间为 [pivot + 1, right]
		} else {
			quickSort(nums, pivot+1, right) // 递归排序右子数组
			right = pivot - 1               // 剩余未排序区间为 [left, pivot - 1]
		}
	}
}

// 取中位数
func medianThree(nums []int, left, mid, right int) int {
	if (nums[left] < nums[mid]) != (nums[left] < nums[right]) {
		return left
	} else if (nums[mid] < nums[left]) != (nums[mid] < nums[right]) {
		return mid
	}
	return right
}

// 哨兵划分
func partition(nums []int, left, right int) int {
	// 找中位数，基准数优化
	mid := medianThree(nums, left, (left+right)/2, right)
	// 中位数交换至最左端，nums[left]作为基准数
	nums[left], nums[mid] = nums[mid], nums[left]
	i, j := left, right
	// 左右交替元素
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 将基准数交换到两字数组的分界线
	nums[i], nums[left] = nums[left], nums[i]
	return i
}
