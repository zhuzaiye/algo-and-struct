// @File:    twoSum_test
// @Version: 1.0.0
// @Creator: JoeLang
// @Date:    2020/8/5 6:42

package two_sum

import (
	"log"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	arr1 := []int{2, 2, 7, 11, 2, 5}
	target1 := 4
	arr2 := []int{2, 7, 11, 5}
	target2 := 9
	log.Println(TwoSum1(arr1, target1))
	log.Println(TwoSum1(arr2, target2))
}
