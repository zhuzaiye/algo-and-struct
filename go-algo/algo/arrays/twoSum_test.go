package arrays

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSumByHash(nums, target))
}
