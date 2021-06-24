package main

import "fmt"

func main() {

	s := make([]string, 5)
	fmt.Println("slice: ", s)
	s[0] = "lex"
	s[1] = "zieme"
	s[2] = "iggy"

	fmt.Println("set: ", s)

	num := []int{1, 2, 3, 4, 5}
	fmt.Println("number", num)

}
