package simulation

// 思路：区分左右有序区间
func findMin(nums []int) int {
	i, j := 0, len(nums)-1
	// 中点和右边界值比
	for i < j {
		m := i + (j-i)/2
		if nums[m] > nums[j] {
			// 左边界收缩到m+1
			i = m + 1
		} else if nums[m] < nums[j] {
			// 右边界收缩到m
			j = m
		} else {
			// 无法判断
			j--
		}
	}
	return nums[i]
}
