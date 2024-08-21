package arrays

import (
	"fmt"
	"testing"
)

func TestRemoveElem(t *testing.T) {
	nums := []int{1, 2, 2, 3, 2, 4}
	length := removeElement(nums, 2)
	fmt.Println(length)

	ll := removeElement2(nums, 2)
	fmt.Println(ll)
}
