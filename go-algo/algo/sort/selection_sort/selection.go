package selection_sort

/**
	选择排序：
		通过两层嵌套的循环，比较数组中元素。
		选择排序内层循环不交换，而是选择最小值的位置，外层循环交换最小值添加到头部，最终实现按照最小值排序。

	属于位置排序，效率差, 平均复杂度和最坏复杂度是O(n**2)
 */

//从小到大的选择排序
func SelectionSort(arr []int) []int {
	for i:=0; i<len(arr); i++{
		minIndex := i
		for j:=i+1; j<len(arr); j++{
			if arr[j] > arr[minIndex]{ //> 从大到小排序；< 从小到大排序
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
	return arr
}

