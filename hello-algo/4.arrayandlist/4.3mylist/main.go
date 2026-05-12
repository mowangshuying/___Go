package main

import (
	"fmt"
)

type MyList struct {
	arrCap int
	arr []int
	arrSize int
	extendRatio int
}

func newMyList() *MyList {
	return &MyList {
		arrCap:10,
		arr: make([]int, 10),
		arrSize:0,
		extendRatio:2,
	}
}

func (this *MyList) size() int {
	return this.arrSize
}

func (this *MyList) capacity() int {
	return this.arrCap
}

func (this *MyList) get(index int) int {
	if index < 0 || index >= this.arrSize {
		panic("index out of range")
	}
	
	return this.arr[index]
}

func (this *MyList) set(index int, num int) {
	this.arr[index] = num
}

func (this *MyList) add(num int) {
	if this.arrSize == this.arrCap {
		this.extendCapacity()
	}
	
	this.arr[this.arrSize] = num
	this.arrSize++
}

func (this *MyList) insert(index int, num int) {
	if index < 0 || index >= this.arrSize {
		panic("index out of range")
	}
	
	if this.arrSize == this.arrCap {
		this.extendCapacity()
	}
	
	for j := this.arrSize - 1; j >= index; j-- {
		this.arr[j+1] = this.arr[j]
	}
	
	this.arrSize++
}

func (this *MyList) remove(index int) int {
	if index < 0 || index >= this.arrSize {
		panic("index out of range")
	}
	
	num := this.arr[index]
	for j := index; j < this.arrSize - 1; j++ {
		this.arr[j] = this.arr[j+1]
	}
	this.arrSize--
	return num
}



func (this *MyList) extendCapacity() {
	this.arr = append(this.arr, make([]int, this.arrCap * (this.extendRatio - 1))...)
	this.arrCap = len(this.arr)	
}

func (this *MyList) toArray() []int {
	return this.arr[:this.arrSize]
}

func main() {
	mylist := newMyList()
	mylist.add(1)
	mylist.add(3)
	mylist.add(5)
	mylist.add(4)
	mylist.add(2)
	
	mylist.remove(3)
	fmt.Printf("list:\n%v", mylist)
}


