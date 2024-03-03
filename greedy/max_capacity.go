package greedy

import "math"

/**
 * 输入一个数组 ht ，其中的每个元素代表一个垂直隔板的高度。数组中的任意两个隔板，以及它们之间的空间可以组成一个容器。
 * 容器的容量等于高度和宽度的乘积（面积），其中高度由较短的隔板决定，宽度是两个隔板的数组索引之差。
 * 请在数组中选择两个隔板，使得组成的容器的容量最大，返回最大容量
 */

func maxCapacity(ht []int) int {
	i, j := 0, len(ht)-1
	res := 0
	for i < j {
		capacity := int(math.Min(float64(ht[i]), float64(ht[j]))) * (j - i)
		res = int(math.Max(float64(res), float64(capacity)))
		if ht[i] < ht[j] {
			i++
		} else {
			j--
		}
	}
	return res
}
