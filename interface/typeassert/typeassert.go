package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (self *Square) Area() float32 {
	return self.side * self.side
}

type Circle struct {
	radius float32
}

func (self *Circle) Area() float32 {
	return self.radius * self.radius * math.Pi
}

func main() {
	var intf Shaper

	square := new(Square)
	square.side = 5

	intf = square

	// t, ok := intf.(*Square)
	// if ok {
	// 	fmt.Printf("The type of intf:%T\n", t)
	// }

	t, ok := intf.(*Circle)
	if ok {
		fmt.Printf("The type of intf:%T\n", t)
	} else {
		fmt.Printf("intf does not contain variable of type Cicle\n")
	}
}
