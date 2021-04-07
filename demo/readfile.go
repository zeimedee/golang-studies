package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("hi.go")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(file)
}
