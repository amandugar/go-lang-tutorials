package main

import "fmt"

const LoginToken string = "ghhghghhg" // Captial Indicates Public

func main() {
	var username string = "aman"
	fmt.Println(username)
	fmt.Printf("Varibale is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Varibale is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Varibale is of type: %T \n", smallVal)

	// var smallFloat float32 = 255.455888487545878774
	var smallFloat float64 = 255.455888487545878774
	fmt.Println(smallFloat)
	fmt.Printf("Varibale is of type: %T \n", smallFloat)

	// default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Varibale is of type: %T \n", anotherVariable)

	// implicit type

	var website = "kahanchale.com"
	fmt.Println(website)

	// no var style
	// not allowded outside main

	numberOfUsers := 300000
	fmt.Println(numberOfUsers)

	// fmt.Println("Variables")
}
