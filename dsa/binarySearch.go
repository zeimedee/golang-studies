package main

import "fmt"

func binary(array []int, k int, l int) bool {
	var n bool
	for i := range array {
		if i == l {
			n = true
		}
	}
	return n
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := len(array) / 2

	b1 := binary(array, n, 5)

	b2 := binary(array, 0, 10)

	if b1 && b2 == true {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
