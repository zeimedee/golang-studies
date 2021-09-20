// calculate average of five integers
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var array [5]int
	var sum int = 0
	var num_string string
	var avg float32

	for i := 0; i < 5; i++ {
		fmt.Println("\nEnter an number")
		// input number
		fmt.Scanln(&num_string)
		// convert to integer
		num, err := strconv.Atoi(num_string)
		// assign num to array index
		array[i] = num
		// assign err to blank identifier
		_ = err
		// calculate sum
		sum += array[i]
	}

	avg = float32(sum) / 5.0
	fmt.Printf("\narray:%v\nsum:%d\naverage:%f", array, sum, avg)
}
