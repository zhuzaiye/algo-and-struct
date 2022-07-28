package commonfix

func LongestCommonPrefix(strs []string) string {
	if len(strs) < 1 { //strs == nil 和 len(strs) < 0 不能对 [] 进行判断 问题1
		return ""
	}
	prefix := strs[0] //问题1的出错在这里？如何细品slice

	// 内外循环为什么不能换？ 逻辑哪里出错了！！
	for j := 0; j < len(prefix); j++ {
		for i := 1; i < len(strs); i++ {
			if j == len(strs[i]) || strs[i][j] != prefix[j] {
				return prefix[:j]
			}
		}
	}
	return prefix
}

/**
 */