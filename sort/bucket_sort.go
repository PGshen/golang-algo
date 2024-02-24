package sort

import (
	"math"
	"sort"
)

// 桶排序
func bucketSort(nums []int) []int {
	min, max := getMinMax(nums)
	bucketCnt := len(nums) / 2
	// 桶间距, 不能用整除，因为需要向上取整
	interval := math.Ceil(float64(max-min+1) / float64(bucketCnt))
	buckets := make([][]int, bucketCnt)
	for i := 0; i < bucketCnt; i++ {
		buckets[i] = make([]int, 0)
	}
	// 分桶
	for _, num := range nums {
		index := int(float64(num-min) / interval)
		buckets[index] = append(buckets[index], num)
	}
	// 桶内排序
	for i := 0; i < bucketCnt; i++ {
		if len(buckets[i]) > 10 {
			// 桶内元素过多，继续分桶
			buckets[i] = bucketSort(buckets[i])
		} else {
			// insertionSort(buckets[i])
			// 使用内置排序算法
			sort.Ints(buckets[i])
		}
	}
	// 合并桶
	sortedNums := make([]int, 0, len(nums))
	for i := 0; i < len(buckets); i++ {
		sortedNums = append(sortedNums, buckets[i]...)
	}
	return sortedNums
}

// 获取最大最小值
func getMinMax(nums []int) (int, int) {
	min := nums[0]
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return min, max
}
