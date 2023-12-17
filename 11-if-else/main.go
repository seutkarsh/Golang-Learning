package main

import "fmt"

func main() {
	fmt.Println("Welcome to If Else in Go lang")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Less than 10"
	} else if loginCount > 10 {
		result = "Greater than 10"
	} else {
		result = "exactly 10"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("Number is Odd")
	}

	// declear a variable and check at the same time

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}

}
