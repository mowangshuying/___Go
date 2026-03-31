package main

import "fmt"

func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("i = %d", i) }

		g(i)

		fmt.Printf(" - g type = %T, g value = %v\n", g, g)
	}
}

func main() {
	f()
}
