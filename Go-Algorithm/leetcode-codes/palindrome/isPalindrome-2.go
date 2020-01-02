package palindrome

/**
时间复杂度：O(log10(n))
空间复杂度：O(1)
*/
func IsPalindrome2(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) { //筛除小于零和不是零但个位数为零的数字
		return false
	}
	revertNumber := 0
	for x > revertNumber {
		revertNumber = revertNumber*10 + x%10 //折叠比较
		x /= 10
	}
	return x == revertNumber || x == revertNumber/10 //偶数位，奇数位数字的比较
}
