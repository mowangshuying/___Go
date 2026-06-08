package main

import "fmt"

type ListNode struct {
	_val int
	_next *ListNode
}


func newListNode (val int) *ListNode {
	return &ListNode{_val:val, _next: nil}
}


func insertNode(n0 *ListNode, p *ListNode) {
	n0next := n0._next
	n0._next = p
	p._next = n0next
}

func removeNode(n0 *ListNode) {
	if n0._next == nil {
		return
	}
	
	n0next := n0._next
	n0nextnext := n0next._next
	n0._next = n0nextnext
}

func acessNode(list *ListNode, index int) *ListNode {
	// if list == nil {
		// return nil
	// }
	
	res := list
	for i:= 0; i < index; i++ {
		if list == nil {
			return nil
		}
		
		res = res._next
	}
	return res
}

func findNode(list *ListNode, target int) *ListNode {
	node := list
	for node != nil {
		if (node._val == target) {
			break
		}
		node = node._next
	}
	return node
}

func main() {
	n0 := newListNode(1)
	n1 := newListNode(3)
	n2 := newListNode(2)
	n3 := newListNode(5)
	n4 := newListNode(4)
	
	n0._next = n1
	n1._next = n2
	n2._next = n3
	n3._next = n4
	
	n := findNode(n0, 4)
	fmt.Printf("n:%d\n", n._val)
}