// Package strings: 给你两个二进制字符串，返回它们的和（用二进制表示

package strings

import (
	"strconv"
	"strings"
)

// char '0' ASCII = 48
// char '1' ASCII = 49

func addBinary(a, b string) string {
	// 如果二者有一个为空返回空
	if strings.Trim(a, " ") == "" || strings.Trim(b, " ") == "" {
		return ""
	}
	ans := ""
	carryVal := 0
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		sum := carryVal
		if i >= 0 {
			sum += int(a[i] - '0')
		}
		if j >= 0 {
			sum += int(b[j] - '0')
		}
		ans = strconv.Itoa(sum%2) + ans
		carryVal = sum / 2
	}
	if carryVal > 0 {
		ans = "1" + ans
	}
	return ans
}
