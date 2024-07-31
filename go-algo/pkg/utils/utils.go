// Date: 2020/11/28 14:05
// Site: HeFei AnHui China
// Auth: JoeLang
// DESC: 定义公共方法

package utils

import "fmt"

// CreateLNode 创建链表,按顺序创建链表
func CreateLNode(node *LNode, max int) {
	currentNode := node
	for i := 1; i < max; i++ {
		currentNode.Next = &LNode{Data: i}
		//currentNode.Next.Data = i
		currentNode = currentNode.Next
	}
}

// PrintLNode 打印链表
func PrintLNode(info string, node *LNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}

func Abs(x int) int {
	switch {
	case x < 0:
		return -x
	case x == 0:
		return 0
	}
	return x
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinBetweenTree(a, b, c int) int {
	min := Min(a, b)
	return Min(min, c)
}

func MaxBetweenTree(a, b, c int) int {
	max := Max(a, b)
	return Max(max, c)
}
