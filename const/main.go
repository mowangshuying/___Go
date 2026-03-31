package main

import "fmt"

func main() {
	const Pi = 3.1415926
	const b1 string = "abc" // explicit type
	const b2 = "abc"        // implicit type

	const beef, two, c = "eat", 2, "veg"
	// const Moday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6
	const (
		Moday, Thursday, Wednesday = 1, 2, 3
		Thursday, Friday, Saturday = 4, 5, 6
	)
	
	// enum
	const (
		Unknown = 0,
		Female = 1,
		Male = 2
	)
}
