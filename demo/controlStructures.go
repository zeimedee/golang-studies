package main

import (
	"fmt"
	"strconv"
)

func main() {
	// if-else
	nums := 20
	if nums >= 10 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	var numString string
	fmt.Println("check if a number is positive or negative")
	fmt.Printf("\n enter a number\n")
	fmt.Scanln(&numString)
	num, err := strconv.Atoi(numString)
	_ = err
	if num > 0 {
		fmt.Printf("\n%d is positive\n", num)
	} else if num < 0 {
		fmt.Printf("\n%d is negative\n", num)
	} else {
		fmt.Printf("\n%d is zero\n", num)
	}

	// switch
	option := 3

	switch option {
	case 1:
		fmt.Println("\noption 1\n")
	case 2:
		fmt.Println("\noption 2\n")
	case 3:
		fmt.Println("\noption 3\n")
	default:
		fmt.Println("\n no match\n")
	}
}
