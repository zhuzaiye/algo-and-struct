/**
 概览：
	 快速排序的核心 -> 选出一个基本点，通过一次遍历得出比它小的A段和比它大的B段的无序数组，
	 然后通过递归方式分别操作A与B，并将低值A结果放在基点左边，高值B结果放在基点右边。

	 全过程都是为了“基准点” pivot 归位。
	 # runtime
	 - worst: O(n^2)
	 - Best: O(n*lgn)
	 - Average: O(n*lgn)
*/

package quick_sort

/**
 简单快排：
	返回小于，等于，大于基点元素的三个数组，基点元素是数组中间值，递归完后将三个数组合并成排序后结果。
 性能：
	数据量“过百”后性能就很差。
 关键点：
	基点比较元素(pivot)
	低值数组
	等值数组
	高值数组
*/
func SimpleQuickSort(numsList []int) []int {
	//判断序列的长度，若长度不大于1，直接输出
	if len(numsList) <= 1 {
		return numsList
	}

	//----------递归模板区---------------
	pivot := numsList[len(numsList)/2] //pivot选择序列中间数值
	partitionFunc := func(arr []int, pivot int) ([]int, []int, []int) {
		lessArr := make([]int, 0)
		greatArr := make([]int, 0)
		equalArr := make([]int, 0)
		for _, v := range arr {
			switch {
			case v < pivot:
				lessArr = append(lessArr, v)
			case v > pivot:
				greatArr = append(greatArr, v)
			default:
				equalArr = append(equalArr, v)
			}
		}
		return lessArr, equalArr, greatArr
	}
	less, equal, greater := partitionFunc(numsList, pivot)
	//-----------------------------------

	copy(numsList, SimpleQuickSort(less))
	copy(numsList[len(less):], equal)
	copy(numsList[len(less)+len(equal):], SimpleQuickSort(greater))

	return numsList
}

/**
 Lomuto快排：
 	以最右边为基点比较元素，递归模板不反悔排序完的数组，而是返回遍历后基点比较元素应该所在的位置下表P
	分为左边小于基点的低值数组，右边是高于基点的高值数组
 性能：
	比霍尔快排简单、性能稍差
 关键点：
	基点比较元素
	遍历索引
	低值数组锚点
*/
func LomutoQuickSort(list []int, low, high int) {
	if low >= high {
		return
	}
	//-----------递归主程序------------
	pivot := list[high]
	i := low
	for j := low; j < high; j++ {
		if list[j] <= pivot {
			list[i], list[j] = list[j], list[i]
			i++
		}
	}
	list[i], list[high] = list[high], list[i]
	//--------------------------------
	LomutoQuickSort(list, low, i-1)
	LomutoQuickSort(list, i+1, high)
}

/**
 霍尔快排：
	递归模板同老毛桃快排，返回低值区的锚点下标。不同的是霍尔快排以最左为基点比较元素，
	并从两端同时进行比较双向排序, 将双方都需调整位置的两个元素相互交换直到双方相交，
	然后将最左元素与右端当前锚点交换，最终实现基准点归位。
 性能：

 关键点：

 */
func HoareQuickSort(list []int, low, high int) {
	if low >= high {
		return
	}
	//------------------递归主程序---------------------
	pivot := list[low]
	right := high
	left := low
	for{
		for list[right] >= pivot && right > low {
			right --
		}
		for list[left] <= pivot && high > left {
			left ++
		}
		if left < right {
			list[left], list[right] = list[right], list[left]
		}else {
			break
		}
	}
	list[low], list[right] = list[right], list[low]
	//--------------------------------------------------
	HoareQuickSort(list, low, right)
	HoareQuickSort(list, right+1, high)
}