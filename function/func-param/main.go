package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func callfunc(y int, f func(int, int) int) {
	res := f(y, y)
	fmt.Printf("res=%d\n", res)
}

func main() {
	callfunc(5, add)
}
