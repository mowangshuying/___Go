package main

import (
	"fmt"
	"strings"
)

func main() {
	// \n \r  \t \u \\

	// str := "beginning of the string" + "second part of the string"
	// s := "hel" + "lo,"
	// s += "world"

	// fmt.Println(s)

	// Prefix & Suffix
	var str string = "This is an example of string"
	fmt.Printf("is Prefix:%t\n", strings.HasPrefix(str, "This"))

	// Contains(s, substr)
	// Index
	// LastIndex
	// Replace (str, old, new, n)
	// Count(s, str)
	// ToLower
	// ToUpper
	// Repeat

	manyG := "gg&gg&g"
	fmt.Printf("mangG has %d <gg>\n", strings.Count(manyG, "gg"))
}
