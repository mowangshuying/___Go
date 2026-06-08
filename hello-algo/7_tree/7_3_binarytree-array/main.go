package main

import "fmt"

// 二叉树的数组表示
// 1.若某节点的索引为i, 则该节点的左节点为2*i+1, 右节点为2*i+2
type ArrayBinaryTree struct {
	_tree []any
}

func newArrayBinaryTree(arr []any) *ArrayBinaryTree {
	return &ArrayBinaryTree{
		_tree: arr,
	}
}

func (abt *ArrayBinaryTree) size() int {
	return len(abt._tree)
}

func (abt *ArrayBinaryTree) val(i int) any {
	if i < 0 || i >= abt.size() {
		return nil
	}

	return abt._tree[i]
}

func (abt *ArrayBinaryTree) left(i int) int {
	return 2*i + 1
}

func (abt *ArrayBinaryTree) right(i int) int {
	return 2*i + 2
}

func (abt *ArrayBinaryTree) parent(i int) int {
	return (i - 1) / 2
}

func (abt *ArrayBinaryTree) dfs(i int, res *[]any) {
	if abt.val(i) == nil {
		return
	}

	(*res) = append((*res), abt.val(i))
	abt.dfs(abt.left(i), res)
	abt.dfs(abt.right(i), res)
}

func main() {
	// 		n1
	// 		/\
	//	   n2 n3
	//	   /\
	//	  n4 n5
	_tree := []string{"n1", "n2", "n3", "n4", "n5"}
	_anytree := make([]any, len(_tree))
	for i := 0; i < len(_tree); i++ {
		// _anytree.append(_tree[i])
		_anytree[i] = _tree[i]
	}

	_btree := newArrayBinaryTree(_anytree)
	var _res []any
	_btree.dfs(0, &_res)
	for i := 0; i < len(_res); i++ {
		fmt.Printf("%s,", _res[i])
	}
	fmt.Printf("\n")
}
