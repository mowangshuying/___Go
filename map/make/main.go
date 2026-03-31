package main

import "fmt"

func main() {
	var mapLit map[string]int
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapMake := make(map[string]float32)
	mapAssigned = mapLit

	mapMake["key1"] = 4.5
	mapMake["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Printf("mapLit[\"one\"]=%d\n", mapLit["one"])
	fmt.Printf("mapMake[\"key2\"]=%f\n", mapMake["key2"])
	fmt.Printf("mapAssigned[\"two\"]=%d\n", mapAssigned["two"])
	fmt.Printf("mapLit[\"ten\"]=%d\n", mapLit["ten"])
}
