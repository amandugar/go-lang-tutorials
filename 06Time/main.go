package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of Go Lang")
	presentTime := time.Now()

	// Compulsory Format Cant be changed
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	createdDate := time.Date(2020, time.June, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}
