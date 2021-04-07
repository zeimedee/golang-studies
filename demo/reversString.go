package main

import (
	"fmt"
	"strings"
)

func reverseArray(arr []string) []string {
	var result []string
	length := len(arr)
	count := len(arr)
	for i := 0; i < length; i++ {
		result = append(result, arr[count-1])
		count--
	}
	return result
}

func main() {

	array := []string{"a", "l", "e", "x", "a"}

	rev := reverseArray(array)
	word := strings.Join(rev, "")
	fmt.Printf(" reversed array: %v\n", rev)
	fmt.Printf("joined array %s", word)

}
