package dyna_plan

func minCostClimbStair(cost []int) int {
	length := len(cost)
	a, b := cost[0], cost[1]
	for i := 2; i < length; i++ {
		ans := min(a, b) + cost[i]
		a, b = b, ans
	}
	return min(a, b)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
