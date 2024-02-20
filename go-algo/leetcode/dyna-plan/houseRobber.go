package dyna_plan

// House Robber
// 不也能抢劫相邻的房子, 求解能够抢到的最多钱财?

// 1.不抢第i房子 2.抢第i房子
// 状态转移方程 dp[i] = max{dp[i-1], dp[i-2]+nums[i]}

// 1.从记忆搜索开始
// 2.状态定义
// 3.状态推导方程
// 4.初始值定义

func houseRobberDp(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i <= n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[n]
}

func houseRobberDpWithMinSpace(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	dp1, dp2, curr := 0, 0, 0
	for i := 0; i < n; i++ {
		curr = max(dp1, dp2+nums[i])
		dp2 = dp1
		dp1 = curr
	}
	return curr
}
