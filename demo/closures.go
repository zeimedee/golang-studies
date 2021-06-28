package main

import "fmt"

func intSec() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	nextSec := intSec()
	fmt.Println(nextSec())
	fmt.Println(nextSec())
	fmt.Println(nextSec())
	fmt.Println(nextSec())
	nextInt := intSec()
	fmt.Println(nextInt())
}
