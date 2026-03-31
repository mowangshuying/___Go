package main

import "fmt"

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("The i=%d\n", i)
	// }

	// i := 5
	// for i > 0 {
	// 	i = i - 1
	// 	fmt.Printf("The i=%d\n", i)
	// }

	// for range
	str := "Go is a beatiful language!"
	for pos, ch := range str {
		fmt.Printf("Pos=%d, Ch=%c\n", pos, ch)
	}
}
