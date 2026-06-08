package main


import (
	"container/list"
	"fmt"
)

type LinkedListDeque struct {
	_list *list.List
}

func newLinkedListDeque() *LinkedListDeque {
	return &LinkedListDeque {
		_list:list.New(),
	}
}

func (s *LinkedListDeque) pushFirst(value any) {
	s._list.PushFront(value)
}

func (s *LinkedListDeque) pushLast (value any) {
	s._list.PushBack(value)
}

func (s *LinkedListDeque) popFirst() any {
	if s.isEmtpy() {
		return nil
	}
	
	e := s._list.Front()
	s._list.Remove(e)
	return e.Value
}

func (s *LinkedListDeque) popLast() any {
	if s.isEmtpy() {
		return nil
	}
	
	e := s._list.Back()
	s._list.Remove(e)
	return e.Value
}

func (s *LinkedListDeque) peekFirst() any {
	if s.isEmtpy() {
		return nil
	}
	
	e := s._list.Front()
	return e.Value
}

func (s *LinkedListDeque) peekLast() any {
	if s.isEmtpy() {
		return nil	
	}
	
	e := s._list.Front()
	return e.Value
}

func (s *LinkedListDeque) size() int {
	return s._list.Len()
}

func (s *LinkedListDeque) isEmtpy()  bool {
	return s._list.Len() == 0
}

func (s *LinkedListDeque) toList() *list.List {
	return s._list
}

func (s *LinkedListDeque) logself() {
	for e := s._list.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d,", e.Value)
	}
	fmt.Printf("\n")
}


func main() {
	q := newLinkedListDeque()
	q.pushFirst(1)
	q.pushFirst(3)
	q.pushFirst(5)
	q.pushFirst(4)
	q.pushLast(2)
	
	q.logself()
	
	q.popFirst()
	q.popLast()
	q.logself()
} 

