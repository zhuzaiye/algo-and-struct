// @File:    twoSum
// @Version: 1.0.0
// @Creator: JoeLang
// @Date:    2020/8/5 6:40

package two_sum

func twoSum(nums []int, target int) []int {
	data := make(map[int]([]int), 0)
	for i,v := range nums {
		_, ok := data[v]
		if ok {
			data[v] = append(data[v], i)
		}else if !ok {
			data[v] = []int{i}
		}
	}
	res := make([]int, 0)
	for value, cal := range data {
		rest := target - value
		if rest == value && len(cal) >= 2 {
			if cal[0] < cal[1] {
				res = append(res, cal[0])
				res = append(res, cal[1])
			}else{
				res = append(res, cal[1])
				res = append(res, cal[0])
			}
			break
		}else if val, ok := data[rest]; ok && rest != value {
			if cal[0] < val[0]{
				res = append(res, cal[0])
				res = append(res, val[0])
			}else {
				res = append(res, val[0])
				res = append(res, cal[0])
			}
			break
		}
	}
	return res
}

func TwoSum1(nums []int, target int) []int {
	rst := make([]int, 0)
	couple := make(map[int]int)
	loc := make(map[int]int)
	for index, value := range nums {
		complete := target - value
		if _, ok := couple[complete]; ok {
			rst = append(rst, loc[complete], index)
			return rst
		}else {
			couple[value] = complete
		}
		loc[value] = index
	}
	return rst
}
