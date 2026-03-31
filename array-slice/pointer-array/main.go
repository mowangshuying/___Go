package main

import "fmt"

func fp(a *[3]int) {
	fmt.Printf("%d\n", a)
}

func main() {
	for i := 0; i < 3; i++ {
		fp(&[3]int{i, i * i, i * i * i})
	}
}
