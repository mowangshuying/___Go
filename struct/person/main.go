package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstname string
	lastname  string
}

func UpperPerson(p *Person) {
	p.firstname = strings.ToUpper(p.firstname)
	p.lastname = strings.ToUpper(p.lastname)
}

func main() {
	p := new(Person)
	p.firstname = "chris"
	p.lastname = "woodward"
	UpperPerson(p)
	fmt.Printf("firstname=%s, lastname=%s\n", p.firstname, p.lastname)
}
