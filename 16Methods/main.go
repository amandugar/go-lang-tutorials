package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")
	// no Inheritance in go lang; No super or parent

	aman := User{"Aman Dugar", "aman@go.dev", true, 20}
	fmt.Println(aman)
	fmt.Printf("Details of Aman are %+v\n", aman)
	fmt.Printf("Name is %v Email is %v. \n", aman.Name, aman.Email)
	aman.GetStatus()
	aman.NewMail()
	fmt.Printf("Name is %v Email is %v. \n", aman.Name, aman.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user Active ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is ", u.Email)
}
