package main

import "fmt"

func main() {
	i := 0
_label:
	fmt.Printf("i=%d\n", i)
	i++
	if i == 5 {
		return
	}

	goto _label
}
