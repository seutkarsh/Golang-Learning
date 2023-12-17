package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to Our Program")

	//to get input we need a buffer reader, this is done using package called bufio...

	reader := bufio.NewReader(os.Stdin) //NewReader function takes an argument to take input from where..this is where os library is used

	fmt.Println("Enter Your Name")

	//comma ok || comma err syntax

	input, _ := reader.ReadString('\n') //the fn takes argument defining when to stop taking input  -- in this case when \n is hit

	fmt.Println("Input entered: ", input)

}
