// Date: 2020/11/28 13:52
// Site: HeFei AnHui China
// Auth: JoeLang
// DESC: golang 语言下的数据结构和对应方法

package utils

//链表定义
type LNode struct {
	Data interface{}
	Next *LNode
}

func NewLNode() *LNode {
	return &LNode{}
}

//二叉树定义
type BNode struct {
	Data       interface{}
	LeftChild  *BNode
	RightChild *BNode
}

func NewBNode() *BNode {
	return &BNode{}
}

//Trie树
type TrieNode struct {
	IsLeaf bool
	Url    string
	Child  []*TrieNode
}

func NewTrieNode(count int) *TrieNode {
	return &TrieNode{
		IsLeaf: false,
		Url:    "",
		Child:  make([]*TrieNode, count),
	}
}

//AVL树
type AVLNode struct {
	Data   int
	Height int
	Count  int
	Left   *AVLNode
	Right  *AVLNode
}

func NewAVLNode(data int) *AVLNode {
	return &AVLNode{
		Data:   data,
		Height: 1,
		Count:  1,
		Left:   nil,
		Right:  nil,
	}
}
