// DATE: 2022/1/16 11:13
// DESC: link list 的基础
// LinkAddr: https://github.com/isdamir/gotype/blob/master/base.go

package linklist

import (
	"fmt"
)

// LinkNode 链表定义
type LinkNode struct {
	Data interface{}
	Next *LinkNode
}

func NewLinkNode() *LinkNode {
	return &LinkNode{}
}

// AddNode 创建链表
func AddNode(node *LinkNode, data interface{}) {
	node.Next = NewLinkNode()
	node.Next.Data = data
}

// CreateLink 将数组转换成链表
func CreateLink(node *LinkNode, arr ...interface{}) {
	curr := node
	for _, v := range arr {
		AddNode(curr, v)
		curr = curr.Next
	}
}

// PrintLink 打印链表
func PrintLink(node *LinkNode) {
	for curr := node.Next; curr != nil; curr = curr.Next {
		fmt.Print(curr.Data, " -> ")
	}
	fmt.Println()
}

func ReverseLink(node *LinkNode) {
	if node == nil || node.Next == nil {
		return
	}
	var next *LinkNode
	var pre *LinkNode
	curr := node.Next // 从第一个Node点开始
	for curr != nil {
		next = curr.Next // 存储当节点后续
		curr.Next = pre  // 反转当前节点

		pre = curr     // 当前节点提升为前节点
		curr = next    // 后续节点提升为当前节点
		PrintLink(pre) // 打印每次反转序列
	}
	node.Next = pre // 续接原来的头节点
	PrintLink(node)
}

func ReverseLinkInsert(node *LinkNode) {
	if node == nil || node.Next == nil {
		return
	}
	var next, curr *LinkNode // 存储节点变量
	curr = node.Next.Next    // 截断第一个节点
	node.Next.Next = nil     // 第一个节点反转为最后节点
	for curr != nil {
		next = curr.Next      // 存储当前节点的后续节点
		curr.Next = node.Next // 反转当前节点
		node.Next = curr      // 作为node序列的第一个节点
		curr = next           // 提升后续节点作为当前节点
	}
}

func RemoveLinkDup(node *LinkNode) {
	if node == nil || node.Next == nil {
		return
	}
	dupMap := make(map[interface{}]struct{})
	for node.Next != nil {
		if _, ok := dupMap[node.Next.Data]; ok {
			node.Next = node.Next.Next
		} else {
			dupMap[node.Data] = struct{}{}
			node = node.Next
		}
	}
}
