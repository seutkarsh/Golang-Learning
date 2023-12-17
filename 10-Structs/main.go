package main

import "fmt"

func main() {

	fmt.Println("Welcome to Structs")

	//IMP: NO Inheritance in GO Lang: No Super or Class or Parent

	//creating an instance of struct
	utkarsh := User{"Utkarsh", "utkarsh@gmail.com", true, 18}

	fmt.Println(utkarsh) // {Utkarsh utkarsh@gmail.com true 18}

	//for representation with keyname use %+v keyword
	fmt.Printf("Visual Representation: %+v\n", utkarsh)

	//for accessing variable value
	fmt.Println("Value of Name:", utkarsh.Name)

}

//declearing a struct  -- type Name struct

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
