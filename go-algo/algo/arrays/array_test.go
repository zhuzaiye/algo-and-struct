// Date: 2020/11/29 9:22
// Site: HeFei AnHui China
// Auth: JoeLang
// DESC:

package arrays

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	//FindDuoByHashDemo()
	FindDuoByXORDemo()
}

func FindDuoByHashDemo() {
	arr := []int{1,3,4,2,5,3}
	fmt.Println("Hash法")
	fmt.Println(FindDupByHash(arr))
}

func FindDuoByXORDemo() {
	arr := []int{1,3,4,2,5,3}
	fmt.Println("Hash法")
	fmt.Println(FindDupByXOR(arr))
}


