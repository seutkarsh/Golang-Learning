package main

import "fmt"

const LoginToken string = "ljabndkabd" // here const means value cant change over time  ---   First letter capital makes the variable "public" hence accessable by anything in this pacakge

func main() {

	// variable name shows error if not used anywhere in file
	//format print is done using printf    --- %T is for type of variable

	//type: string
	var username string = "utkarsh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	//type: bool
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	//type: int -- we use int as in common where memory is taken care of automatically..we can also use types like uint8, unit32 etc...
	var value int = 256
	fmt.Println(value)
	fmt.Printf("Variable is of type: %T \n", value)

	//type: float -- use types like float32, float64....for precision
	var floatValue float64 = 256.2659415
	fmt.Println(floatValue)
	fmt.Printf("Variable is of type: %T \n", floatValue)

	//default values and aliases

	var anotherValue int //only declearing the variable - go provides 0 as default value to any initialized int
	fmt.Println(anotherValue)
	fmt.Printf("Variable is of type: %T \n", anotherValue)

	var anotherString string //only declearing the variable - go provides empty string as default value to any initialized string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type: %T \n", anotherString)

	//implicit style
	//here lexer decides the type and make it work
	var website = "Hello.com"
	fmt.Println(website)

	//no var style  --- only aloowed inside a method...DO NOT USE IN GLOBAL CONTEXT

	numberOfChocolates := 3500 // the colon then equals is syntax that shows its a variable assignment with decleartion without using var keyword
	fmt.Println(numberOfChocolates)

}
