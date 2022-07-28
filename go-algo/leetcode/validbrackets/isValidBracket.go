package validbrackets

/*  注意：这里的顺序必须是()[]{},但是([)]就不行，但是辅助栈的思想在后面很有用
	问题：只包含从(), [], {}的字符串，判断字符串是否有效？
		1、左括号必须用相同类型右括号闭合；
		2、左括号必须以正确顺序闭合
		3、空字符串是有效字符

	注意：顺序和成对出现
	算法分析：
		- 遇到左括号，入栈； 遇到右括号， 出栈
	提前返回false：
		- 发现不符合就返回，提升算法效率
		- 边界问题，栈为空和符号序列以左括号结尾
*/
func IsValidBracket(s string) bool {
	if s == "" {
		return true
	}
	bracketDict := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"!": "!",
	}
	bracketStack := []string{"!"}
	var x string
	for _, c := range s {
		if _, ok := bracketDict[string(c)]; ok {
			bracketStack = append(bracketStack, string(c))
		} else if x, bracketStack = bracketStack[len(bracketStack)-1], bracketStack[:len(bracketStack)-1]; bracketDict[x] != string(c) {
			return false
		}
	}
	return len(bracketStack) == 1
}
