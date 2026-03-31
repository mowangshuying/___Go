package main

import "fmt"

func sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

func main() {
	array := [3]float64{7.0, 8.5, 9.1}

	x := sum(&array)
	fmt.Printf("The Sum of the array is:%f\n", x)
}
