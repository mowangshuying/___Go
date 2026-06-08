package binarysearchtree

import "fmt"

type TreeNode struct {
	_val   int
	_left  *TreeNode
	_right *TreeNode
}

func NewTreeNode(num int) *TreeNode {
	return &TreeNode{
		_left:  nil,
		_right: nil,
		_val:   num,
	}
}

type BinarySearchTree struct {
	_root *TreeNode
}

func NewBinarySearchTree() *BinarySearchTree {
	bst := &BinarySearchTree{}
	bst._root = nil
	return bst
}

func (bst *BinarySearchTree) GetRoot() *TreeNode {
	return bst._root
}

func (bst *BinarySearchTree) Insert(num int) {
	cur := bst.GetRoot()
	if cur == nil {
		bst._root = NewTreeNode(num)
		return
	}

	var pre *TreeNode = nil

	for cur != nil {
		if cur._val == num {
			return
		}
		pre = cur

		if cur._val < num {
			cur = cur._right
		} else {
			cur = cur._left
		}
	}

	node := NewTreeNode(num)
	if pre._val < num {
		pre._right = node
	} else {
		pre._left = node
	}
}

func (bst *BinarySearchTree) Remove(num int) {
	cur := bst.GetRoot()
	if cur == nil {
		return
	}

	var pre *TreeNode = nil
	for cur != nil {
		if cur._val == num {
			break
		}

		pre = cur
		if cur._val < num {
			cur = cur._right
		} else {
			cur = cur._left
		}
	}

	if cur == nil {
		return
	}

	if cur._left == nil || cur._right == nil {
		var child *TreeNode = nil
		if cur._left != nil {
			child = cur._left
		} else {
			child = cur._right
		}

		if cur != bst._root {
			if pre._left == cur {
				pre._left = child
			} else {
				pre._right = child
			}

		} else {
			bst._root = child
		}
	} else {
		/// 子节点数量为2
		tmp := cur._right
		for tmp._left != nil {
			tmp = tmp._left
		}

		bst.Remove(tmp._val)
		cur._val = tmp._val
	}
}

func (bst *BinarySearchTree) LogAnum(node *TreeNode) {
	// cur := bst.getRoot()
	if node != nil {
		// fmt.Printf("%d,", cur._val)
		bst.LogAnum(node._left)
		fmt.Printf("%d,", node._val)
		bst.LogAnum(node._right)
	}
}

func (bst *BinarySearchTree) Logself() {
	fmt.Printf("bst:")
	bst.LogAnum(bst.GetRoot())
	fmt.Printf("\n")
}
