package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Functions")

	greeter()
	result := adder(3, 5)

	fmt.Println("Value is:", result)

	proResult := proAdder(1, 2, 3, 4, 5)

	fmt.Println("Pro Adder Value is:", proResult)

}

func greeter() {
	fmt.Println("Hello")
}

// func nameOfFunction(paramName paramType, paramName paramtype) returnTypeIfAny{}
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// to recieve non defined no of params but of same type
func proAdder(values ...int) int {
	total := 0

	for _, value := range values {
		total = total + value
	}

	return total
}

// if we need to return two values   --- capture result using comma ok syntax
// func proAdderNew(values ...int) (int,string) {
// 	total := 0

// 	for _, value := range values {
// 		total = total + value
// 	}

// 	return total,"Hello"
// }
