package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("\n Reverse an array")
	var arr [5]int
	var num_str string

	for i := 0; i < 5; i++ {
		fmt.Println("\n Enter a number")
		fmt.Scanln(&num_str)
		num, err := strconv.Atoi(num_str)
		_ = err
		arr[i] = num
	}
	fmt.Printf("\n array:%v", arr)
	// reverse array
	for i := 0; i < 5/2; i++ {
		temp := arr[i]
		arr[i] = arr[4-i]
		arr[4-i] = temp
	}
	fmt.Printf("\n Reversed array: %v", arr)
}
