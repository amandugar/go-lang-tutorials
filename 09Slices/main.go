package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices in go Lang")

	var fruitList = []string{"Apple", "Tomato", "Peaches"}
	fmt.Printf("Type of Fruit List is: %T", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 2345
	highScores[2] = 341
	highScores[3] = 242
	// highScores[4] = 242 // Will Give Error
	highScores = append(highScores, 555, 666, 777)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

}
