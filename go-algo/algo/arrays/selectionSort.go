// 选择排序

package arrays

// 找到arr中最小项
func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestLocation := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestLocation = i
		}
	}
	return smallestLocation
}

// 选择排序
func selectionSort(arr []int) []int {
	ret := make([]int, 0, len(arr))
	length := len(arr)
	for i := 0; i < length; i++ {
		smallestIndex := findSmallest(arr)
		ret = append(ret, arr[smallestIndex])
		arr = append(arr[:smallestIndex], arr[smallestIndex+1:]...)
	}
	return ret
}
