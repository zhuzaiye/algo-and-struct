package arrays

// 二分法搜索
// 定位
func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

// 二分法搜索
func searchBinary(nums []int, target int) int {
	ans := -1
	left, right := 0, len(nums)-1 // 1 首尾
	for left <= right {           // 2
		mid := (right-left)>>1 + left
		if target < nums[mid] {
			right = mid - 1 // 3
		} else if target > nums[mid] {
			left = mid + 1 // 4
		} else {
			ans = mid
			break
		}
	}
	return ans
}

func searchRange(nums []int, target int) []int {
	posRange := []int{-1, -1}
	posLeft := searchInsert(nums, target)
	posRight := searchInsert(nums, target+1)
	if posLeft >= len(nums) || nums[posLeft] != target {
		return posRange
	}
	return []int{posLeft, posRight - 1}
}
