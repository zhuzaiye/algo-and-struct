package arrays

// 给定一个`整数数组 nums` 和一个`整数目标值 target`
// 在该数组中找出 `和为目标值 target` 的那 两个 整数，并返回它们的`数组下标`

// 暴力枚举
// 时间复杂度：O(N^2) 空间复杂度：O(1)

func twoSum01(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[i] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// HashMap的方法
// O(1) 内存占用大 主要是提前预设内存
func twoSumByHash(nums []int, target int) []int {
	indexMap := make(map[int][]int, len(nums))
	for i := range nums {
		if _, ok := indexMap[nums[i]]; ok {
			indexMap[nums[i]] = append(indexMap[nums[i]], i)
		} else {
			indexMap[nums[i]] = []int{i}
		}
	}
	for i, num := range nums {
		if v, ok := indexMap[target-num]; ok {
			if (target - num) == num {
				if len(v) > 1 {
					return []int{i, v[1]}
				}
				continue
			}
			return []int{i, v[0]}
		}
	}
	return nil
}

func twoSumByHash2(nums []int, target int) []int {
	indexMap := make(map[int]int, len(nums))
	for i, x := range nums {
		if loc, ok := indexMap[target-x]; ok {
			return []int{loc, i}
		}
		indexMap[x] = i
	}
	return nil
}
