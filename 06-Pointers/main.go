package main

import "fmt"

func main() {

	fmt.Println("Welcome to Pointers!!!")

	// initializing a pointer named ptr, astrick shows its a pointer storing int value
	// var ptr *int

	// fmt.Printf("Type of variable ptr, %T \n", ptr)
	// fmt.Println("Value of ptr, ", ptr)

	myNumber := 23

	ptr := &myNumber

	fmt.Println("Value of pointer, ", ptr)       // prints 0xc00000a0b8   -- some memory location
	fmt.Println("Value of pointer, ", *ptr)      // prints 23
	fmt.Println("Value of pointer, ", &myNumber) //prints memory location

	// "&" provides the memory location of the variable  provided a variable  -- reference
	// "*" provides the value strored in the momeory location provided a memory location  --value

	*ptr = *ptr + 2 //changes the value of myNumber tp 25

	fmt.Println("New Value: ", myNumber)
}
