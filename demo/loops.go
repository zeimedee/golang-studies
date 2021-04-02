package main

import (
	"fmt"
)

func main() {
	fmt.Println("loops")
	count := 0
	for count <= 10 {
		fmt.Printf("\n count is %d", count)
		count++
	}

	fmt.Println("\n short declared for loop")

	for count := 1; count <= 10; count++ {
		fmt.Printf("%d\n", count*5)
	}
}
