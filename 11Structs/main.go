package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")
	// no Inheritance in go lang; No super or parent

	aman := User{"Aman Dugar", "aman@go.dev", true, 20}
	fmt.Println(aman)
	fmt.Printf("Details of Aman are %+v\n", aman)
	fmt.Printf("Name is %v Email is %v.", aman.Name, aman.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
