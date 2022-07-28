package search

// 二叉树搜索, 递归方法
func binarySearchByRecursive(arr []int, target, lowIndex, highIndex int) int {
	if highIndex < lowIndex || len(arr) == 0 {
		return -1
	}
	midIndex := (highIndex + lowIndex) / 2
	if arr[midIndex] > target {
		return binarySearchByRecursive(arr, target, lowIndex, midIndex-1)
	} else if arr[midIndex] < target {
		return binarySearchByRecursive(arr, target, midIndex+1, highIndex)
	} else {
		return midIndex
	}
}

// 循环二叉树搜索
func binarySearchByIter(arr []int, target, lowIndex, highIndex int) int {
	startIndex := lowIndex
	endIndex := highIndex
	length := len(arr)
	if length == 0 || highIndex > length || lowIndex < 0 {
		return -1
	}
	var mid int
	for startIndex < endIndex {
		mid = (startIndex + endIndex) / 2
		if arr[mid] > target {
			endIndex = mid - 1
		} else if arr[mid] > target {
			startIndex = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
