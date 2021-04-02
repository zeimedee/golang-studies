package main

import "fmt"

func main() {
	var firstname, lastname, city string = "alex", "zieme", "accra"
	var age int = 26
	fmt.Printf("\nFirstname: %s \nLastname: %s \nAge: %d \nCity: %s \n", firstname, lastname, age, city)

	var str string
	fmt.Println("\nEnter something")
	fmt.Scanln(&str)
	fmt.Printf("something: %s\n", str)

	var str1 string
	fmt.Println("Enter something else:")
	fmt.Scanf("%s", &str1)
	fmt.Printf("something else: %s", str1)
}
