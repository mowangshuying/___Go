package main

import "fmt"

func main() {
	var arr [5]int

	for i := 0; i < len(arr); i++ {
		arr[i] = i * 2
	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d]=%d\n", i, arr[i])
	}
}
