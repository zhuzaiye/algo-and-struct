package dyna_plan

// 爬楼梯

/*
	回溯法: 站在上一次结果上, 尝试下一次, 直到最后结果, 如果最后结果不满足, 回到前一次, 尝试下一个方案
*/

func backtrack(choices []int, state, n int, res []int) {
	// 每升到一个台阶n, 方案数就加一
	if state == n {
		res[0] += 1
	}
	// 遍历所有选择
	for _, choice := range choices {
		// 剪枝: 不允许超过第n阶
		if state+choice > n {
			continue
		}
		// 尝试: 选择,更新状态
		backtrack(choices, state+choice, n, res)
		// 回退
	}
}

func climbStairsBacktrack(n int) int {
	// 爬楼方式, 1阶或2阶
	choices := []int{1, 2}
	// 从0阶开始爬
	state := 0
	// 方案存储结构
	res := make([]int, 1)
	res[0] = 0
	backtrack(choices, state, n, res)
	return res[0]
}

// dp[i] = dp[i-1]+dp[i-2] 深度优先搜索
// 暴力搜索 递归树的深度n 复杂度: O(2^n)
// 问题在于: 子问题重叠
// 思考解决: 如果子问题只被计算一次? 重复计算进行剪枝

func dfs(i int) int {
	if i == 1 || i == 2 {
		return i
	}
	count := dfs(i-1) + dfs(i-2)
	return count
}

func climbStairsDFS(n int) int {
	return dfs(n)
}

// 记忆搜索: 复杂度降低到 O(n)
// 1. 首次计算dp[i], mem[i]存储记忆
// 2. 再次计算需要dp[i]时, 不在进行递归计算, 直接从mem中去获取, 进行重叠子递归计算剪枝

func dfsMem(i int, mem []int) int {
	if i == 1 || i == 2 {
		return i
	}
	// 剪枝: 如果dp[i]已经记忆, 则直接返回
	if mem[i] != -1 {
		return mem[i]
	}
	count := dfsMem(i-1, mem) + dfsMem(i-2, mem)
	// 首次计算, 记忆dp[i]
	mem[i] = count
	return count
}

func climbStairsDFSMem(n int) int {
	mem := make([]int, n+1)
	for i := range mem {
		mem[i] = -1
	}
	return dfsMem(n, mem)
}

// 以上都是从上至下的拆解问题
// 动态规划从下到上的拆解问题, 其不包含回溯过程, 是需要循环迭代实现. 不需要递归

/*
	总体思考方式: 归纳总结,由n-1和n-2的结果, 推理出n的逻辑, 将问题拆解成更小问题, 通过存储子问题避免重复计算.
    动态规划五步法
	1. 数据存储， 三个变量
	2. 第n阶是是从第n-1阶跨1步或者从第n-2阶跨2步，所以跨上第n阶的方法数是n-1阶和n-2阶跨上方法的和
	3. 初始化 1阶 2阶 n阶
	4. 遍历顺序
*/

func climbStairsDP(n int) int {
	if n < 3 {
		return n
	}
	// 从初始状态开始, 进行至下往上的状态转移, 直到目标状态时结束
	// dp表, 对应分解步骤的状态值
	dp := make([]int, n+1)
	// todo: 使用滚动变量替换dp数组, 降低空间复杂度
	dp[1] = 1 // 最小子问题, 初始状态
	dp[2] = 2 // 最小子问题, 初始状态

	for i := 3; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2] // 状态转移方程
	}
	return dp[n]
}

// 无后效性: 约束性动态规划需要扩展状态定义, 是的问题诚信满足无后效性
// 其定义为：给定一个确定的状态，它的未来发展只与当前状态有关，而与过去经历的所有状态无关。
// 问题: 给定一个共有 n 阶的楼梯，你每步可以上1阶或者2阶，但不能连续两轮跳1阶，请问有多少种方案可以爬到楼顶？

// dp[i,1] = dp[i-1, 2]
// dp[i,2] = dp[i-2,1] + dp[i-2,2]

/* 带约束爬楼梯：动态规划 */
func climbingStairsConstraintDP(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	// 初始化 dp 表，用于存储子问题的解
	dp := make([][3]int, n+1)
	// 初始状态：预设最小子问题的解
	dp[1][1] = 1
	dp[1][2] = 0
	dp[2][1] = 0
	dp[2][2] = 1
	// 状态转移：从较小子问题逐步求解较大子问题
	for i := 3; i <= n; i++ {
		dp[i][1] = dp[i-1][2]
		dp[i][2] = dp[i-2][1] + dp[i-2][2]
	}
	return dp[n][1] + dp[n][2]
}

// 步骤：描述决策，定义状态，建立dp表，推导状态转移方程，确定边界条件等。
// 子问题分解是一种从顶至底的思想，因此按照“暴力搜索->记忆化搜索->动态规划”的顺序实现更加符合思维习惯
