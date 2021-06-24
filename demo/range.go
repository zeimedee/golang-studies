package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	items := map[string]int{"a": 2, "b": 3, "c": 4}

	for k, v := range items {
		fmt.Printf("%s -> %d\n", k, v)
	}
}
