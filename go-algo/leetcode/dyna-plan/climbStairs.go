package dyna_plan

// 爬楼梯

/*
	动态规划五步法
	1. 数据存储， 三个变量
	2. 第n阶是是从第n-1阶跨1步或者从第n-2阶跨2步，所以跨上第n阶的方法数是n-1阶和n-2阶跨上方法的和
	3. 初始化 1阶 2阶 n阶
	4. 遍历顺序
*/
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	a, b, ans := 1, 2, 0
	for i := 2; i < n; i++ {
		ans = a + b
		a, b = b, ans
	}
	return ans
}

func climbStairs1(n int) int {
	if n < 3 {
		return n
	}
	return climbStairs1(n-1) + climbStairs1(n-2)
}
