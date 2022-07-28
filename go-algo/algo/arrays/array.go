// Date: 2020/9/14 23:37
// DESC: golang 的数组实例操作与纪要

package arrays

// FindDupByHash find the duplicated item by hash
//时间复杂度: O(n)
//空间复杂度: O(n)
func FindDupByHash(arr []int) []int {
	if arr == nil {
		return nil
	}
	dup := make([]int, 0, len(arr))
	hashPool := make(map[int]struct{})
	for _, elem := range arr {
		if _, ok := hashPool[elem]; ok {
			dup = append(dup, elem)
		} else {
			hashPool[elem] = struct{}{}
		}
	}
	return dup
}

// FindDupByXOR xor calc
func FindDupByXOR(arr []int) int {
	if arr == nil {
		return -1
	}
	rst := 0
	for _, v := range arr {
		rst ^= v
	}
	l := len(arr)
	for i := 1; i < l; i++ {
		rst ^= i
	}
	return rst
}
