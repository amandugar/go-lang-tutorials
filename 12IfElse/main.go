package main

import (
	"fmt"
)

func main() {
	fmt.Println("If else in Go Lang")
	var result string
	loginCount := 23

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "Exactlty 10"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is Not less than 10")
	}
}
