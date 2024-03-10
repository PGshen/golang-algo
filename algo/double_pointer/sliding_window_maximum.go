package doublepointer

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	windows := []int{}
	for i := 0; i < k; i++ {
		push(&windows, nums[i])
	}
	res = append(res, front(windows))
	for i := k; i < len(nums); i++ {
		pop(&windows, nums[i-k])
		push(&windows, nums[i])
		res = append(res, front(windows))
	}
	return res
}

// 实现单调递减队列
func push(windows *[]int, val int) {
	// 从后往前，踢出比val小的值，保证单调递减
	for len(*windows) > 0 && (*windows)[len(*windows)-1] < val {
		*windows = (*windows)[:len(*windows)-1]
	}
	*windows = append(*windows, val)
}

func pop(windows *[]int, val int) {
	// 如果相等就出队列，相当意味着不在窗口内了
	if len(*windows) > 0 && (*windows)[0] == val {
		*windows = (*windows)[1:]
	}
}

func front(windows []int) int {
	return windows[0]
}
