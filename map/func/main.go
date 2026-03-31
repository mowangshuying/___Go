package main

import "fmt"

// map 初始化後面要加 ","
func main() {
	mapFunc := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 30 },
	}

	fmt.Printf("mapFunc[1]() = %d\n", mapFunc[1]())
}
