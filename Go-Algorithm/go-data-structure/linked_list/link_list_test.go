// Date: 2020/11/28 15:10
// Site: HeFei AnHui China
// Auth: JoeLang
// DESC: 链表测试

package linked_list

import (
	"Go-Algorithm/pkg/utils"
	"fmt"
	"testing"
)

func TestLinkListReverse(t *testing.T) {
	//linkListReverseMode1()
	//linkListReverseMode2()
	linkListReverseMode3()
}

func linkListReverseMode1() {
	head := &utils.LNode{}
	fmt.Println("就地逆序：")
	utils.CreateLNode(head, 8)
	utils.PrintLNode("逆序前：", head)
	ReverseInPlace(head)
	utils.PrintLNode("逆序后：", head)
}

func linkListReverseMode2() {
	head := &utils.LNode{}
	fmt.Println("就地逆序：")
	utils.CreateLNode(head, 8)
	utils.PrintLNode("逆序前：", head)
	RecursiveReverse(head)
	utils.PrintLNode("逆序后：", head)
}

func linkListReverseMode3() {
	head := &utils.LNode{}
	fmt.Println("就地逆序：")
	utils.CreateLNode(head, 8)
	utils.PrintLNode("逆序前：", head)
	InsertReverse(head)
	utils.PrintLNode("逆序后：", head)
}