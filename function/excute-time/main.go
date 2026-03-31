package main

import (
	"fmt"
	"time"
)

func dosomethings() {
	time.Sleep(5 * time.Second)
}

func main() {
	s := time.Now()
	dosomethings()
	e := time.Now()
	d := e.Sub(s)
	fmt.Printf("take time: %s\n", d)
}
