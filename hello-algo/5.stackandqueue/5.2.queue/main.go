package main

import (
	"container/list"
	"fmt"
)

type LinkedListQueue struct {
	_list *list.List
}

func newLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue {
		_list: list.New(),
	}
}

func (s *LinkedListQueue) push(value any) {
	s._list.PushBack(value)
} 

func (s *LinkedListQueue) pop() any {
	if s.isEmpty() {
		return nil
	}
	
	e := s._list.Front()
	s._list.Remove(e)
	return e.Value
}

func (s *LinkedListQueue) peek() any {
	if s.isEmpty() {
		return nil
	}
	
	e := s._list.Front()
	return e.Value
}

func (s *LinkedListQueue) size() int {
	return s._list.Len()
}

func (s *LinkedListQueue) isEmpty() bool {
	return s._list.Len() == 0
}

func (s *LinkedListQueue) toList() *list.List {
	return s._list
}

func (s *LinkedListQueue) logself() {
	for e := s._list.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d,", e.Value)
	}
	fmt.Printf("\n")
}

func main() {
	q := newLinkedListQueue()
	q.push(1)
	q.push(3)
	q.push(5)
	q.push(4)
	q.push(2)
	
	q.logself()
	
	q.pop()
	q.pop()
	q.logself()
}