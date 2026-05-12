package main

import (
	"container/list"
	"fmt"
)

type LinkedListStack struct {
	_list *list.List
}

func newLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		_list:list.New(),
	}
}

func (s *LinkedListStack) push(value int) {
	s._list.PushBack(value)
}

func (s *LinkedListStack) pop() any {
	if s.isEmpty() {
		return nil
	}
	
	e := s._list.Back()
	s._list.Remove(e)
	return e.Value
}

func (s *LinkedListStack) peek() any {
	if s.isEmpty() {
		return nil
	}
	
	e := s._list.Back()
	return e.Value
}

func (s *LinkedListStack) size() int {
	return s._list.Len()
}


func (s *LinkedListStack) isEmpty() bool {
	return s._list.Len() == 0
}

func (s *LinkedListStack) toList() *list.List {
	return s._list
}

func (s *LinkedListStack) logself() {
	for e := s._list.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d,", e.Value)
	}
	fmt.Printf("\n")
}

func main() {
	/// 
	s := newLinkedListStack()
	s.push(1)
	s.push(3)
	s.push(5)
	s.push(4)
	s.push(2)
	
	s.logself()
	
	s.pop()
	s.pop()
	s.logself()
	
	fmt.Printf("peek:%d\n", s.peek())
}
