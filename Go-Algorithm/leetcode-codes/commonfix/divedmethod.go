package commonfix

/**
 分治：
	LCP(S1···Sn) = LCP(LCP(S1,···,Sk),LCP(Sk+1,···,Sn))
*/

func longestCommonPrefix(strs []string, l int, r int) string {
	if l == r {
		return strs[1]
	} else {
		mid := (l + r) / 2
		lcpLeft := longestCommonPrefix(strs, 1, mid)
		lcpRight := longestCommonPrefix(strs, mid+1, r)
		return commonPrefix(lcpLeft, lcpRight)
	}
}

func commonPrefix(left, right string) string {
	var min int
	if len(left) < len(right) {
		min = len(left)
	} else {
		min = len(right)
	}
	for i := 0; i < min; i++ {
		if left[i] != right[i] {
			return left[:i]
		}
	}
	return left[:min]
}

func LongestCommonPrefix2(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	return longestCommonPrefix(strs, 0, len(strs)-1)
}
