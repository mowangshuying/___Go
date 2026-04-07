package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//demo0
	// fmt.Printf("Please enter your full name:")

	// var firstname string
	// var lastname string
	// fmt.Scanln(&firstname, &lastname)
	// fmt.Printf("Hi %s %s\n", firstname, lastname)

	// demo1
	// fmt.Printf("Please enter some input:\n")
	// inputR := bufio.NewReader(os.Stdin)
	// input, err := inputR.ReadString('\n')
	//
	//	if err == nil {
	//		fmt.Printf("The input is: %s\n", input)
	//	}

	// demo2
	inputR := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter your name:\n")
	input, err := inputR.ReadString('\n')
	if err != nil {
		fmt.Printf("There were errors reading, exiting program\n")
	}

	fmt.Printf("You name is %s", input)

	switch input {
	case "Philip\r\n":
		fmt.Printf("welcome Philip!\n")
	case "Chris\r\n":
		fmt.Printf("Welcome Chris!\n")
	case "Ivo\r\n":
		fmt.Printf("Welcome Ivo!\n")
	}
}
