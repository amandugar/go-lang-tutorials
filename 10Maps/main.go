package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all the Languages: ", languages)
	fmt.Println("JS shorts for ", languages["JS"])
	delete(languages, "RB")
	fmt.Println("List of all the Languages: ", languages)

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
