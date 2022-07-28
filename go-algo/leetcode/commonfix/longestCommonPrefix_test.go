package commonfix

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	testData1 := []string{"flowers", "flow", "flight"}
	testData3 := []string{"flowers", "", "flight"}
	testData2 := make([]string, 0)
	rst1 := LongestCommonPrefix(testData1)
	fmt.Println(rst1)
	rst2 := LongestCommonPrefix(testData2)
	fmt.Println(rst2)
	rst3 := LongestCommonPrefix(testData3)
	fmt.Println(rst3)
}