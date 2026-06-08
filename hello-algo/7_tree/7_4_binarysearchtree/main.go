package main

import (
	// "container/list"
	BST "___Go/hello-algo/7_tree/7_4_binarysearchtree/binarysearchtree"
	// "fmt"
)

func main() {
	bst := BST.NewBinarySearchTree()
	bst.Insert(7)
	bst.Insert(4)
	bst.Insert(5)
	bst.Insert(8)
	bst.Logself()
}
