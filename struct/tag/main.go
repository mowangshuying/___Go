package main

import (
	"fmt"
	"reflect"
)

type TagType struct {
	filed1 bool   "an important answer"
	filed2 string "The name of the thing"
	filed3 int    "how much there are"
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}

func main() {
	tt := TagType{true, "Barak obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}
