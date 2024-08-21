package arrays

import "testing"

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	ans := searchInsert(nums, 2)
	println(ans)
}
