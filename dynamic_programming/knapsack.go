package dynamicprogramming

import "math"

// 给定 n 个物品，第 i 个物品的重量为 wgt[i-1]、价值为val[i-1] ，和一个容量为 cap  的背包。
// 每个物品只能选择一次，问在限定背包容量下能放入物品的最大价值。
// 注意：物品编号i从1开始，数组索引从0开始计数

// 暴力搜索
// c 为背包剩余容量
func knapsackDFS(wgt, val []int, i, c int) int {
	if i == 0 || c == 0 {
		return 0
	}
	// 超过了剩余背包容量，只能选择不放入背包
	if wgt[i-1] > c {
		return knapsackDFS(wgt, val, i-1, c)
	}
	// 计算不放入和放入背包两种情况的最大价值
	no := knapsackDFS(wgt, val, i-1, c)
	yes := knapsackDFS(wgt, val, i-1, c-wgt[i-1]) + val[i-1]
	// 返回价值更大的方案
	return int(math.Max(float64(no), float64(yes)))
}

// 记忆搜索
func knapsackDFSMem(wgt, val []int, mem [][]int, i, c int) int {
	if i == 0 || c == 0 {
		return 0
	}
	if mem[i][c] != -1 {
		return mem[i][c]
	}
	// 超过了剩余背包容量，只能选择不放入背包
	if wgt[i-1] > c {
		return knapsackDFSMem(wgt, val, mem, i-1, c)
	}
	// 计算不放入和放入背包两种情况的最大价值
	no := knapsackDFSMem(wgt, val, mem, i-1, c)
	yes := knapsackDFSMem(wgt, val, mem, i-1, c-wgt[i-1]) + val[i-1]
	// 返回价值更大的方案
	mem[i][c] = int(math.Max(float64(no), float64(yes)))
	return mem[i][c]
}

// 动态规划
func knapsackDP(wgt, val []int, c int) int {
	n := len(wgt)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, c+1)
	}
	// 状态转移
	for i := 1; i <= n; i++ {
		for j := 1; j <= c; j++ {
			if wgt[i-1] > j {
				// 剩余重量放不下，只能不选
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i-1][j-wgt[i-1]]+val[i-1])))
			}
		}
	}
	return dp[n][c]
}

// 动态规划，空间优化
/* 0-1 背包：空间优化后的动态规划 */
func knapsackDPComp(wgt, val []int, cap int) int {
	n := len(wgt)
	// 初始化 dp 表
	dp := make([]int, cap+1)
	// 状态转移
	for i := 1; i <= n; i++ {
		// 倒序遍历，这里只能倒叙，否则会被覆盖
		for c := cap; c >= 1; c-- {
			if wgt[i-1] <= c {
				// 不选和选物品 i 这两种方案的较大值
				dp[c] = int(math.Max(float64(dp[c]), float64(dp[c-wgt[i-1]]+val[i-1])))
			}
		}
	}
	return dp[cap]
}
