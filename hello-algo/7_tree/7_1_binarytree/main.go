package main

import (
	"fmt"
	"container/list"
)

type TreeNode struct {
	_val int
	_left *TreeNode
	_right *TreeNode
}

func newTreeNode(v int) *TreeNode {
	return &TreeNode {
		_left: nil,
		_right: nil,
		_val: v,
	}
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
	
}