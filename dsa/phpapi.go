package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fetch, err := http.Get("http://localhost:8080/roleauth/api.php")
	if err != nil {
		fmt.Println(err)
	}

	res, err := ioutil.ReadAll(fetch.Body)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(res)
}
