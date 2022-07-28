package insertion_sort

/**
插入排序适合少量或者基本有序的数据，数据量大的时候复杂度很大。
*/
func InsertionSort(arr []int) []int {
	length := len(arr)
	if length == 1 {
		return arr
	}
	for i := 1; i < length; i++ {
		backup := arr[i]
		j := i - 1
		//将选出的被排数比较后插入左边有序区域
		for j >= 0 && backup < arr[j] { //j>=0必须放在前面，否则会越界
			arr[j+1] = arr[i] //移动有序数组
			j-- //反向移动下标
		}
		arr[j+1] = backup //插入移动后的空位
	}
	return arr
}
