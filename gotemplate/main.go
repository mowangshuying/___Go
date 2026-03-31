package main

import "fmt"

const c = "C"

var v int = 5

type T struct{}

func init() { // initialization of package
	fmt.Print("__init\n")
}

func main() {
	// var a int

	Func1()
}

func (t T) Method1() {

}

func Func1() {}
