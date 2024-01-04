package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to Golang web Request")

	response, err := http.Get(url)

	checkNilErr(err)

	defer response.Body.Close() // caller's responsibility to close the connection....ALWAYS CLOSE THE CONNECTION

	//read response, response is in http.response type...read using ioutils

	data, err := ioutil.ReadAll(response.Body)

	checkNilErr(err)

	content := string(data)

	fmt.Println(content)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
