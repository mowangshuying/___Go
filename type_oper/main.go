package main

import "fmt"
import "math"

func intToUint8(n int) (uint8, error)
{
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}


func main() {
	// var avar = 10
	// b := true
	// b = avar != 5
	// b = avar == 10

	// fmt.Printf("b=%t\n", b)

	// int & float
	// var n int16 = 34
	// var m int32

	// m = int32(n)

	// fmt.Printf("32 bit int is:%d\n", m)
	// fmt.Printf("16 bit int is:%d\n", n)

	// %d 整数
	// %g 浮点数
	// %x 16进制数
	// %f  浮点数
	// %e 科学计数
	// %0d
	// %n.mg
	// %5.2e
}
