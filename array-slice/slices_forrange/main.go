package main

import "fmt"

func main() {
	var slice []int = make([]int, 4)

	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	slice[3] = 4

	for i, v := range slice {
		fmt.Printf("i=%d,v=%d\n", i, v)
	}

	seasons := []string{"spring", "summer", "autumn", "winter"}
	for i, v := range seasons {
		fmt.Printf("i=%d,v=%s\n", i, v)
	}
}
