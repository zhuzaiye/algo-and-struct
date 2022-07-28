package search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	array := []int{1, 4, 6, 7, 8, 9, 10}
	//index := binarySearchByRecursive(array, 7, 0, 3)
	index := binarySearchByIter(array, 7, 0, 3)
	fmt.Println(index)
}
