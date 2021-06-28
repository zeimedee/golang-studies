package main

import "fmt"

func add(a, b int) int {
	return a + b
}

//functions with multiple return values

func two() (int, int) {
	return 3, 7
}

//variadic functions are functions that can be called with any number of trailing arguments

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	add := add(2, 2)
	fmt.Printf("sum: %d\n", add)

	a, b := two()
	fmt.Println(a)
	fmt.Println(b)

	v, _ := two()
	fmt.Println(v)
	fmt.Printf("\n variadic function \n")
	num := []int{1, 2, 3, 4, 5}
	sum(num...)
}
