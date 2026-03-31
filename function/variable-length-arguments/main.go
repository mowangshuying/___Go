package main

import "fmt"

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}

	_t := s[0]
	for _, v := range s {
		if v < _t {
			_t = v
		}
	}

	return _t
}

func main() {
	x := min(1, 2, 3, -1)
	fmt.Printf("x=%d\n", x)
}
