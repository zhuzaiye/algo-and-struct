package array

import "fmt"

// 亲前后左右指针
// 左指针记录nums中等于val的位置，
// 右指针进行遍历nums
func removeElement(nums []int, val int) int {
	lp := 0
	for index := range nums {
		if nums[index] != val {
			nums[lp] = nums[index]
			lp++
		}
	}
	fmt.Println(nums[:lp])
	return lp
}


// 首尾指针
func removeElement2(nums []int, val int) int {
	head, tail := 0, len(nums)
	for head < tail {
		if nums[head] == val {
			nums[head] = nums[tail-1]
			tail--
		}else{
			head++
		}
	}
	fmt.Println(nums)
	return head
}