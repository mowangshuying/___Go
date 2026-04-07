package main

import "fmt"

type S struct {
	i int
	f float32
	s string
}

func main() {
	s := new(S)
	s.i = 10
	s.f = 15.5
	s.s = "chris"

	fmt.Printf("s:\n%v\n", s)
}
