package main

import "fmt"

func main() {

	fmt.Println("Welcome to Structs")
	utkarsh := User{"Utkarsh", "utkarsh@gmail.com", true, 18}
	fmt.Println(utkarsh)
	fmt.Printf("Visual Representation: %+v\n", utkarsh)
	fmt.Println("Value of Name:", utkarsh.Name)
	utkarsh.NewEmail()
	fmt.Println(utkarsh)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

//methods are fucntions associated with a struct  -- in go methods are not defined inside struct, a struct is passed into the method.
//IMP: changes to struct will not work since a copy of the original struct is passed in the method...hence use of pointers is expected

func (u User) NewEmail() {
	u.Email = "newEmail@gmail.com"
	fmt.Println("New Email: ", u.Email)
}
