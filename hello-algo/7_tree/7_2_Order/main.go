package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	_val   int
	_left  *TreeNode
	_right *TreeNode
}

func newTreeNode(v int) *TreeNode {
	return &TreeNode{
		_left:  nil,
		_right: nil,
		_val:   v,
	}
}

func levelOrder(root *TreeNode) any {
	queue := list.New()
	queue.PushBack(root)
	nums := make([]any, 0)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode) // get the first node;
		nums = append(nums, node._val)                  // append it

		if node._left != nil {
			queue.PushBack(node._left)
		}

		if node._right != nil {
			queue.PushBack(node._right)
		}
	}
	return nums
}

func preOrder(root *TreeNode) { // 前序遍历
	if root == nil {
		return
	}

	fmt.Printf("%d,", root._val)
	preOrder(root._left)
	preOrder(root._right)
}

func inOrder(root *TreeNode) { // 中序遍历
	if root == nil {
		return
	}
	inOrder(root._left)
	fmt.Printf("%d,", root._val)
	inOrder(root._right)
}

func postOrder(root *TreeNode) { // 后序遍历
	if root == nil {
		return
	}

	postOrder(root._left)
	postOrder(root._right)
	fmt.Printf("%d,", root._val)
}

func main() {
	n1 := newTreeNode(1)
	n2 := newTreeNode(2)
	n3 := newTreeNode(3)
	n4 := newTreeNode(4)
	n5 := newTreeNode(5)

	n1._left = n2
	n1._right = n3
	n2._left = n4
	n2._right = n5

	nums := levelOrder(n1)
	// fmt.Printf("\n")
	fmt.Printf("level order:%v\n", nums)

	fmt.Printf("pre order:")
	preOrder(n1)
	fmt.Printf("\n")

	fmt.Printf("in order:")
	inOrder(n1)
	fmt.Printf("\n")

	fmt.Printf("post order:")
	postOrder(n1)
	fmt.Printf("\n")

}
