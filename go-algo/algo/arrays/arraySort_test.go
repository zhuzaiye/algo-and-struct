package arrays

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr1 := []int{5, 2, 9, 0, 1, 6, 8, 4, 7, 5, 3}
	sortArr1 := selectionSort(arr1)
	fmt.Println("sort arr1: ", sortArr1)

	arr2 := []int{2, 0, 1, 0, 4, 5, 9, 2, 3}
	sortArr2 := selectionSort(arr2)
	fmt.Println("sort arr2: ", sortArr2)
}
