package array

// 一个升序排列的数组nums，原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 元素的相对顺序应该保持一致 。
func removeDuplicates(nums []int) int {
	ans := make([]int, 0, len(nums))
	for i := range nums {
		if i == 0 || nums[i] != ans[len(ans)-1] {
			ans = append(ans, nums[i])
		}
	}
	for i := range ans {
		nums[i] = ans[i]
	}
	return len(ans)
}
