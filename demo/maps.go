package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["n1"] = 3
	m["n2"] = 9

	fmt.Println("map:", m)
	fmt.Println("length", len(m))

	n := map[string]int{"a": 2, "b": 3, "c": 4}
	//check if element exist, map returns 2 items
	k, ok := n["a"]
	if ok {
		fmt.Println(k)
	}

	fmt.Println("n:a", n["a"])
}
