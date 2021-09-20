package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("if array", array, "includes", 4)
	for i := range array {
		if i == 4 {
			fmt.Println(true)
			break
		}
	}
}
