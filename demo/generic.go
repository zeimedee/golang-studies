package main

import "fmt"

func show[T any](s []T) {
	for _, value := range s {
		fmt.Println(value)
	}
}
func main() {

	show([]int{1, 2, 3, 4, 5})
	show([]string{"a", "b", "c", "d", "e"})
}
