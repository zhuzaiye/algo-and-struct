// DATE: 2022/5/11
// DESC:

package binarytree

import "testing"

func TestArrayToBinaryTree(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	binaryTree := ArrayToBinaryTree(arr, 0, len(arr)-1)
	//MidTraverseBinaryTree(binaryTree)
	//PreTraverseBinaryTree(binaryTree)
	LayerTraverseBinaryTree(binaryTree)
}
