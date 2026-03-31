package main

import "fmt"

func func1() {

	fmt.Printf("top\n")

	defer func2()

	fmt.Printf("bottom\n")
}

func func2() {
	fmt.Printf("func2 done.\n")
}

func main() {
	func1()
}
