package main

import "fmt"

func main() {
	var arr [20]int
	var slice []int = arr[2:5]

	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d]=%d\n", i, slice[i])
	}

	fmt.Printf("arr length = %d\n", len(arr))
	fmt.Printf("slice length = %d\n", len(slice))
	fmt.Printf("slice cap=%d\n", cap(slice))

	slice = slice[0:6]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d] = %d\n", i, slice[i])
	}
}
