package listnode

//Building linked list using GoLang
//linked list structure
type ListNode struct {
	Val int
	Next *ListNode
}

func MergeTwoLinkedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil  {
		return l2
	}else if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = MergeTwoLinkedLists(l1.Next, l2)
		return l1
	}else {
		l2.Next = MergeTwoLinkedLists(l1, l2.Next)
		return l2
	}
}