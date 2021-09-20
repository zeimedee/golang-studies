package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {
	file, err := os.ReadFile("hi.go")
	if err != nil {
		log.Fatal(err)
	}
	// os.Stdout.Write(file)
	fmt.Println(reflect.TypeOf(file))
}
