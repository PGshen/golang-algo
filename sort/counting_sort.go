package sort

// 计数排序
func countingSort(nums []int) {
	// 1. 统计数组最大元素 m
	m := 0
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	// 2. 统计各数字的出现次数
	// counter[num] 代表 num 的出现次数
	counter := make([]int, m+1)
	for _, num := range nums {
		counter[num]++
	}
	// 3. 遍历 counter ，将各元素填入原数组 nums
	for i, num := 0, 0; num < m+1; num++ {
		for j := 0; j < counter[num]; j++ {
			nums[i] = num
			i++
		}
	}
}

type People struct {
	name string
	age  int
}

// People对象根据年龄排序
func countingSort2(people []People) {
	// 得到最大年龄，确认计数数组大小
	maxAge := 0
	for _, person := range people {
		if person.age > maxAge {
			maxAge = person.age
		}
	}
	// 用年龄计数
	counter := make([]int, maxAge+1)
	for _, person := range people {
		counter[person.age]++
	}
	// 3. 求 counter 的前缀和，将“出现次数”转换为“尾索引”
	// 即 counter[person.age]-1 是 person 在 res 中最后一次出现的索引
	for i := 0; i < maxAge; i++ {
		counter[i+1] += counter[i]
	}
	// 4. 倒序遍历 people ，将各元素填入结果数组 res
	// 初始化数组 res 用于记录结果
	n := len(people)
	res := make([]People, n)
	for i := n - 1; i >= 0; i-- {
		person := people[i]
		// 将 num 放置到对应索引处
		res[counter[person.age]-1] = person
		// 令前缀和自减 1 ，得到下次放置 num 的索引
		counter[person.age]--
	}
	// 复制结果到people
	copy(people, res)
}
