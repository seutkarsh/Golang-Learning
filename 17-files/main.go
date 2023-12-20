package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to Golang files")
	content := "this is the content for the file"

	file, err := os.Create("./myfile.txt")

	checkNilErr(err)

	lenght, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("Lenght is:", lenght)

	defer file.Close()

	readFile("./myfile.txt")
}

func readFile(fileName string) {
	// data is in format of databyte
	data, err := ioutil.ReadFile(fileName)

	checkNilErr(err)

	fmt.Println("Data in file:", string(data))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
