package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	var s person
	s.firstName = "alex"
	s.lastName = "zieme"
	s.age = 25
	p2 := person{firstName: "zieme", lastName: "axel", age: 23}

	fmt.Printf("\n name: %s %s\n age: %d\n", s.firstName, s.lastName, s.age)
	fmt.Printf("\n name: %s %s\n age: %d\n", p2.firstName, p2.lastName, p2.age)

}
