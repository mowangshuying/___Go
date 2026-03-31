package main

import "fmt"

// test1 --------
func test1() {
	fmt.Printf("In test1 before calling greeting.\n")
	greeting()
	fmt.Printf("In test1 after calling greeting.\n")
}

func greeting() {
	fmt.Printf("In greeting: Hi!\n")
}

// / test2 -------
func test2() {
	fmt.Printf("2*5*6=%d\n", multiply3Nums(2, 5, 6))
}

func multiply3Nums(a int, b int, c int) int {
	return a * b * c
}

// test3 -------
func test3() {
	x2, x3 := getX2AndX3(10)
	fmt.Printf("x2=%d, x3=%d\n", x2, x3)

	_, y3 := getX2AndX3(5)
	fmt.Printf("x2=_, x3=%d\n", y3)
}

func getX2AndX3(intput int) (int, int) {
	return 2 * intput, 3 * intput
}

// test4 ------
func test4() {
	reply := -1
	multiply(1, 1, &reply)
	fmt.Printf("reply=%d\n", reply)
}

func multiply(a, b int, reply *int) {
	(*reply) = a * b
}

func main() {
	// test1()
	// test2()
	// test3()

	test4()
}
