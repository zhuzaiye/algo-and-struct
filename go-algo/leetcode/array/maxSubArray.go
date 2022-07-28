package array

// 问题： 给定数组nums,找到一个具有最大和的连续数组(子数组中最少包含一个数字), 返回这个最大连续数组的和。

// 第一种方法：
// 动态规划为了“找到不同子序列之间的递推关系, 以子序列的结束点为基准”
// 推导公式：sum[i] = max{sum[i-1]+a[i], a[i]}

func maxSubArray2(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	sum := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		// 当前位置到i的前面的子序列的最大值
		sum = func(x, y int) int {
			if x > y {
				return x
			}
			return y
		}(sum+nums[i], nums[i])
		max = func(x, y int) int {
			if x > y {
				return x
			}
			return y
		}(max, sum)
	}
	return max
}

// 该方法会修改输入slice，为保证原有的slice不变和输出对应子序列和子序列值
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ { // 外层遍历从i=1开始进行遍历
		if nums[i]+nums[i-1] > nums[i] { //
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
