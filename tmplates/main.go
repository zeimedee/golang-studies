package main

import (
	"bytes"
	"fmt"
	"html/template"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	t, err := template.ParseFiles("hello.html")
	check(err)

	data := struct{ Name string }{"alexander"}
	var tmp bytes.Buffer
	err = t.Execute(&tmp, data)
	check(err)
	result := tmp.String()
	fmt.Println(result)
}
