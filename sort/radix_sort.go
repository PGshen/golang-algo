package sort

import "math"

// 基数排序
// 排序的正确性依赖于计数排序的稳定性
func radixSort(nums []int) {
	max := math.MinInt
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	// 按照从低位到高位的顺序遍历
	for exp := 1; max >= exp; exp *= 10 {
		countingSortDigit(nums, exp)
	}
}

// 获取num的第k位， 其中exp=10^(k-1)
func digit(num, exp int) int {
	return (num / exp) % 10
}

// 计数排序，按nums的第k位排序
func countingSortDigit(nums []int, exp int) {
	// 10进制，10个桶
	counter := make([]int, 10)
	n := len(nums)
	// 统计第k位 个数字的出现次数
	for i := 0; i < n; i++ {
		d := digit(nums[i], exp)
		counter[d]++
	}
	// 求前缀和，将 出现个数 转换为 数组索引
	for i := 1; i < 10; i++ {
		counter[i] += counter[i-1]
	}
	// 倒序遍历，按照索引填入res
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		d := digit(nums[i], exp)
		// d在数组中的索引
		j := counter[d] - 1
		res[j] = nums[i]
		counter[d]--
	}
	copy(nums, res)
}
