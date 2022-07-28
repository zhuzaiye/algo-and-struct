// DATE: 2022/5/11
// DESC: 二叉树

package binarytree

import (
	"fmt"
	"queue/queue"
)

// BinaryNode 二叉树结点
type BinaryNode struct {
	Value     interface{} // 结点数值
	LeftNode  *BinaryNode // 左-子结点
	RightNode *BinaryNode // 右-子结点
}

// ArrayToBinaryTree 将有序数组转化成完全二叉树 [递归方法]
func ArrayToBinaryTree(arr []int, start, end int) *BinaryNode {
	var root *BinaryNode
	if end >= start {
		mid := (start + end + 1) / 2
		root = &BinaryNode{
			Value:     arr[mid],
			LeftNode:  ArrayToBinaryTree(arr, start, mid-1),
			RightNode: ArrayToBinaryTree(arr, mid+1, end),
		}
	}
	return root
}

// PreTraverseBinaryTree 前序遍历二叉树 根-左-右 [递归方法]
func PreTraverseBinaryTree(root *BinaryNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Value, " ")
	// 遍历左子树
	if root.LeftNode != nil {
		PreTraverseBinaryTree(root.LeftNode)
	}
	if root.RightNode != nil {
		PreTraverseBinaryTree(root.RightNode)
	}
}

// MidTraverseBinaryTree 中序遍历二叉树 左-根-右 [递归方法]
func MidTraverseBinaryTree(root *BinaryNode) {
	if root == nil {
		return
	}
	// 遍历左子树
	if root.LeftNode != nil {
		MidTraverseBinaryTree(root.LeftNode)
	}
	fmt.Print(root.Value, " ")
	if root.RightNode != nil {
		MidTraverseBinaryTree(root.RightNode)
	}
}

func LayerTraverseBinaryTree(root *BinaryNode) {
	if root == nil {
		return
	}
	q := queue.NewArrayQueue()
	var node *BinaryNode
	q.EnQueue(root)
	for q.Size() > 0 {
		node = q.DelQueue().(*BinaryNode)
		fmt.Print(node.Value, " ")
		if node.LeftNode != nil {
			q.EnQueue(node.LeftNode)
		}
		if node.RightNode != nil {
			q.EnQueue(node.RightNode)
		}
	}
}
