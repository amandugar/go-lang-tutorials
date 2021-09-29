package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Array in go Lang")

	var fruitsList [4]string

	fruitsList[0] = "Apple"
	fruitsList[1] = "Tomato"
	fruitsList[3] = "Peach"

	fmt.Println("Fruits List is: ", fruitsList)
	fmt.Println("Fruits List is: ", len(fruitsList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veg List is: ", vegList)
}
