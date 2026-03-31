package main

import "fmt"

func main() {
	var i = 5
	fmt.Printf("An integer:%d, it's location in memory: %p\n", i, &i)

	// nil
	// ptr
	var iptr *int
	iptr = &i

	fmt.Printf("The value at memory laction %p is %d\n", iptr, *iptr)

	s := "good bye"
	var sptr *string = &s

	fmt.Printf("Here is the pointer p: %p\n", sptr)
	fmt.Printf("Here is the string *p: %s\n", *sptr)

	// 不可访问空指针
}
