package main

type ArrayBinaryTree struct {
	_tree []any
}

func newArrayBinaryTree(arr []any) *ArrayBinaryTree {
	return &ArrayBinaryTree {
		_tree:arr,
	}
}

func (abt *ArrayBinaryTree) size() int {
	return len(abt._tree)
}

func (abt *ArrayBinaryTree) val(i int) any{
	if i < 0 || i >= abt.size() {
		return nil
	}
	
	return abt._tree[i]
}

func (abt *ArrayBinaryTree) left(i int) int{
	return 2 * i + 1
}

func (abt *ArrayBinaryTree) right(i int) int {
	return 2 * i + 2
}

func (abt *ArrayBinaryTree) parent(i int) int {
	return (i - 1) / 2
}

func (abt *ArrayBinaryTree) dfs (i int, res *[]any) {
	if abt.val(i) == nil {
		return	
	}
	
	(*res) = append((*res), abt.val(i))
	abt.dfs(abt.left(i), res)
	abt.dfs(abt.right(i), res)
}

func (abt *ArrayBinaryTree) preOrder() []any {
	var res []any
	abt.dfs(0, &res)
	return res
}

func main() {
	
}