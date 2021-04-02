package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputreader := bufio.NewReader(os.Stdin)
	fmt.Printf("\nenter a sentence:\n")
	text, _ := inputreader.ReadString('\n')
	fmt.Printf("sentence is : %s", text)
}
