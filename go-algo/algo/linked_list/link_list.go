// Date: 2020/11/28 14:03
// Site: HeFei AnHui China
// Auth: JoeLang
// DESC: 链表的逆序

package linked_list

import "Go-Algorithm/pkg/utils"

/*
 链表逆序
*/
//带头节点的链表逆序,就地逆序操作
func ReverseInPlace(node *utils.LNode) {
	if node == nil || node.Next == nil { //
		return
	}
	var pre *utils.LNode //前驱节点
	var cur *utils.LNode //当前节点
	next := node.Next    //后继节点

	for next != nil {
		cur = next.Next //将后继节点的后继放到当前
		next.Next = pre //将原来的前驱节点放到后继节点的后继位置
		pre = next
		next = cur
	}
	node.Next = pre
}

//递归式逆序链表
func RecursiveReverseChild(node *utils.LNode) *utils.LNode {
	if node == nil || node.Next == nil {
		return node
	}
	newHead := RecursiveReverseChild(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newHead
}

func RecursiveReverse(node *utils.LNode) {
	firstNode := node.Next
	//递归调用
	newHead := RecursiveReverseChild(firstNode)
	node.Next = newHead
}

//插入式逆序链表
func InsertReverse(node *utils.LNode) {
	if node == nil || node.Next == nil {
		return
	}

	var cur *utils.LNode
	var next *utils.LNode
	cur = node.Next.Next
	//设置链表第一个节点为尾节点
	node.Next.Next = nil
	//逐步将节点插入到头节点
	for cur != nil {
		next = cur.Next
		cur.Next = node.Next
		node.Next = cur
		cur = next
	}
}

/*
 从无序链表中移除重复节点
*/

func RemoveDup(node *utils.LNode) {
	if node == nil || node.Next == nil {
		return
	}

	outerCur := node.Next
	var innerCur *utils.LNode
	var innerPre *utils.LNode
	for ; outerCur != nil; outerCur = outerCur.Next {
		for innerCur, innerPre = outerCur.Next, outerCur; innerCur != nil; {
			if outerCur.Data == innerCur.Data {
				innerPre.Next = innerCur.Next
				innerCur = innerCur.Next
			} else {
				innerPre = innerCur
				innerCur = innerCur.Next
			}
		}
	}
}


