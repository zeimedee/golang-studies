package main

import (
	"fmt"
	"time"
)

func coffee() {
	for i := 0; i < 10; i++ {
		fmt.Println("coffee", i)
		time.Sleep(2 * time.Second)
	}
}

func milk() {
	for i := 0; i < 10; i++ {
		fmt.Println("milk", i)
		time.Sleep(2 * time.Second)

	}
}
func main() {

	go milk()
	coffee()

}
