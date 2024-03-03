package dynamicprogramming

import "math"

// 给定一个 m * n 的二维网格 grid ，网格中的每个单元格包含一个非负整数，表示该单元格的代价。机器人以左上角单元格为起始点，每次只能向下或者向右移动一步，直至到达右下角单元格。请返回从左上角到右下角的最小路径和。

// 解题思路
// 1. 思考每轮的决策，定义状态，从而得到bp表
// 2. 找出最优子结构，进而推导出状态转移方程
// 3. 确定边界条件和状态转移顺序

/**
 * 暴力搜索
 */
func minPathSumDFS(grid [][]int, i, j int) int {
	if i == 0 && j == 0 {
		return grid[0][0]
	}
	if i < 0 || j < 0 {
		return math.MaxInt
	}
	up := minPathSumDFS(grid, i-1, j)
	left := minPathSumDFS(grid, i, j-1)
	return int(math.Min(float64(up), float64(left))) + grid[i][j]
}

/**
 * 记忆化搜索
 */
func minPathSumDFSMem(grid, mem [][]int, i, j int) int {
	if i == 0 && j == 0 {
		return grid[0][0]
	}
	if i < 0 || j < 0 {
		return math.MaxInt
	}
	// 如果曾经计算过，则直接返回
	if mem[i][j] != -1 {
		return mem[i][j]
	}
	up := minPathSumDFSMem(grid, mem, i-1, j)
	left := minPathSumDFSMem(grid, mem, i, j-1)
	mem[i][j] = int(math.Min(float64(up), float64(left))) + grid[i][j]
	return mem[i][j]
}

/**
 * 动态规划
 */
func minPathSumDP(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	dp[0][0] = grid[0][0]
	// 首行
	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	// 首列
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	// 状态转移，其他行、列
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))) + grid[i][j]
		}
	}
	return dp[n-1][m-1]
}

/**
 * 动态规划，空间优化
 */
func minPathSumDPComp(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	// 只需要存储一行
	dp := make([]int, m)
	dp[0] = grid[0][0]
	// 首行
	for j := 1; j < m; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}
	for i := 1; i < n; i++ {
		dp[0] += grid[i][0]
		for j := 1; j < m; j++ {
			// 覆盖更新。min里面的 dp[j-1]是左，dp[j]是上
			dp[j] = int(math.Min(float64(dp[j-1]), float64(dp[j]))) + grid[i][j]
		}
	}
	return dp[m-1]
}
