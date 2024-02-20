package dyna_plan

import "go-algo/pkg/utils"

// 爬楼梯最小消耗[最优方案]

// 动态规划更适合: 最优子结构
/* 爬楼梯最小代价：动态规划 */
func minCostClimbingStairsDP(cost []int) int {
	n := len(cost) - 1
	if n == 1 || n == 2 {
		return cost[n]
	}
	// 初始化 dp 表，用于存储子问题的解
	dp := make([]int, n+1)
	// 初始状态：预设最小子问题的解
	dp[1] = cost[1]
	dp[2] = cost[2]
	// 状态转移：从较小子问题逐步求解较大子问题
	for i := 3; i <= n; i++ {
		dp[i] = utils.Min(dp[i-1], dp[i-2]) + cost[i]
	}
	return dp[n]
}
