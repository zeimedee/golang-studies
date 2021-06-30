package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}
func main() {
	a, err := f1(42)
	if err == nil {
		fmt.Println(a)
	} else {
		fmt.Println(e)
	}
}
