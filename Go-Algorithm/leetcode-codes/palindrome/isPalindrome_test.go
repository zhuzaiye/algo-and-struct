package palindrome

import "testing"

/**
 go测试函数的规则：
	1、go test是go语言的程序测试工具，以*_test.go的文件形式存在，go build并不会将其编译，称为构建的一部分。
	2、测试函数的文件格式必须是funcName_test.go，测试函数必须以Test_funName，传入参数必须是*testing.T
	3、运行测试函数，go test Test_funName，注意特定测试文件必须与源码关联在一起
	4、在测试中看到运行详细结果，可以加上-v参数
*/
func TestIsPalindrome(t *testing.T) {
	testData := map[int]bool{
		121:  true,
		-121: false,
		10:   false,
		11:   true,
	}
	for k, v := range testData {
		actual := IsPalindrome(k)
		if actual != v {
			t.Logf("%d 判断失败！", k)
		}

	}
}
