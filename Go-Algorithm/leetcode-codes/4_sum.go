// Date: 2020/10/13 23:50
// Site: HeFei AnHui China
// Auth: JoeLang
// DESC:

package leetcode_codes

import "sort"

func fourSum(nums []int, target int) [][]int {
	//统计nums中每个数字的数量
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}
	//去重后的nums
	uniqNums := make([]int, 0, len(counter))
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums) //快速排序

	//分类型进行组合计算
	rst := make([][]int, 0)
	for i := 0; i < len(uniqNums); i++ {
		//第一种是4个相同的
		if (uniqNums[i]*4 == target) && counter[uniqNums[i]] >= 4 {
			rst = append(rst, []int{uniqNums[i], uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			//三个相同，一个不同
			if (uniqNums[i]*3 + uniqNums[j] == target) && counter[uniqNums[i]] > 2 {
				rst = append(rst, []int{uniqNums[i], uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if (uniqNums[j]*3 + uniqNums[i] == target) && counter[uniqNums[j]] > 2 {
				rst = append(rst, []int{uniqNums[i], uniqNums[j], uniqNums[j], uniqNums[j]})
			}
			//两个相同+两个相同
			if (uniqNums[i]*2 + uniqNums[j]*2 == target) && counter[uniqNums[j]] > 1 && counter[uniqNums[i]] > 1{
				rst = append(rst, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			for k := j + 1; k < len(uniqNums); k++ {
				//两个相同，一个不同，一个不同
				if (uniqNums[i]*2 + uniqNums[j] + uniqNums[k] == target) && counter[uniqNums[i]] > 1 {
					rst = append(rst, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[k]})
				}
				if (uniqNums[j]*2 + uniqNums[i] + uniqNums[k] == target) && counter[uniqNums[j]] > 1 {
					rst = append(rst, []int{uniqNums[i], uniqNums[j], uniqNums[j], uniqNums[k]})
				}
				if (uniqNums[j] + uniqNums[i] + uniqNums[k]*2 == target) && counter[uniqNums[k]] > 1 {
					rst = append(rst, []int{uniqNums[i], uniqNums[j], uniqNums[k], uniqNums[k]})
				}
				c := target -(uniqNums[i] + uniqNums[j] + uniqNums[k])
				if c > uniqNums[k] && counter[c] > 0 {
					rst = append(rst, []int{uniqNums[i], uniqNums[j], uniqNums[k], c})
				}
			}
		}
	}
	return rst
}
