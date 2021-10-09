package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Go")
	content := "This needs to go in a files - Learn.in"
	file, err := os.Create("./fileTest.txt")

	checkNilErr(err)
	length, err := io.WriteString(file, content)

	checkNilErr(err)
	fmt.Println("length is ", length)
	defer file.Close()
	readFile("./fileTest.txt")
}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println("Text data inside  file is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
