package main

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