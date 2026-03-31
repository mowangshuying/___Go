package main

import "fmt"

func main() {
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	for k, v := range map1 {
		fmt.Printf("k=%d,v=%f\n", k, v)
	}

	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}

	fmt.Printf("items=%v\n", items)
}
