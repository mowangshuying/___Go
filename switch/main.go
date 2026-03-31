package main

import "fmt"

func main() {
	var num int = 100

	// switch num {
	// case 98, 99:
	// 	fmt.Printf("It's equal to 98/99\n")
	// case 100:
	// 	fmt.Printf("It's equal to 100\n")
	// default:
	// 	fmt.Printf("It's not equal every.")
	// }

	switch {
	case num > 0 && num <= 50:
		fmt.Printf("num > 0 && num <= 50\n")
	case num > 50 && num <= 100:
		fmt.Printf("num > 50 && num <= 100\n")
	default:
		fmt.Print("unknow.\n")
	}

	// fallthrough
}
