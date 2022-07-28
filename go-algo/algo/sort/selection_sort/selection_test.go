package selection_sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	nums := []int{3, 2, 8, 1, 9, 4, 7, 6, 0, 5}
	sortNum := SelectionSort(nums)
	fmt.Println(sortNum)
}
