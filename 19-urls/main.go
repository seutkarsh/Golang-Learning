package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?courseId=reactjs&experience=8"

func main() {
	fmt.Println("Welcome to handling URls in golang")

	//parsing url using url library
	result, err := url.Parse(myurl)

	checkNilErr(err)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
