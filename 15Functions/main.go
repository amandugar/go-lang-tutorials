package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions in Go")
	greeter()
	result := adder(3, 5)

	fmt.Println("Result is: ", result)
	proRes, myMessage := proAdder(2, 5, 7, 5, 6, 5, 4, 8)
	fmt.Println(proRes, myMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func greeter() {
	fmt.Println("Namastey from Go Lang")
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "Hi"
}
