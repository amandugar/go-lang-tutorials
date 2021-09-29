package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to a class on pointer")

	// var ptr *int
	// fmt.Println("Value of ptr is: ", ptr)

	myNumber := 23
	var ptr = &myNumber // refrence = &
	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value is: ", myNumber)
}
