package dyna_plan

// 动态规划--斐波那契数

/*
	1. 定义collection存储递归或者动态运算的结果
	2. 确定地推方程式[动态处理逻辑]
	3. collection数据如何初始化[逻辑的开始端的确定]
	4. 遍历顺序[动态逻辑的先后执行顺序]
	5. 演绎步骤【1，2，3，4】 [模拟性的推到演绎确定collection-逻辑-初始化的正确性]
*/

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fib1(n int) int {
	if n < 2 {
		return n
	}
	a, b, c := 0, 1, 0       // 使用变量存储结果，并初始化
	for i := 1; i < n; i++ { // 定义归纳的动态逻辑块和遍历顺序
		c = a + b
		a, b = b, c
	}
	return c
}
