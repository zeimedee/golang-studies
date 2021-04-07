package main

import (
	"fmt"
	"strconv"
)

func reverseArray(ar [5]int) [5]int {
	for i := 0; i < 5/2; i++ {
		temp := ar[i]
		ar[i] = ar[4-i]
		ar[4-i] = temp
	}

	return ar
}

func main() {
	fmt.Println("\n Reverse an array")
	var arr [5]int
	var num_str string
	var reverse [5]int

	for i := 0; i < 5; i++ {
		fmt.Println("\n Enter a number")
		fmt.Scanln(&num_str)
		num, err := strconv.Atoi(num_str)
		_ = err
		arr[i] = num
	}
	fmt.Printf("\n array:%v", arr)

	reverse = reverseArray(arr)
	fmt.Printf("\n reversed array: %v", reverse)
}
