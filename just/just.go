package main

import (
	"___Go/struct/structpack"
	"fmt"
)

func main() {
	s := new(structpack.ExpStruct)
	s.Mi = 5
	s.Mf = 10.5
	fmt.Printf("s:%v\n", s)
}
