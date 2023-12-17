package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays")

	//array declearation - size / length must always be provided
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Mango"

	fmt.Println("Fruit List: ", fruitList)
	fmt.Println("Fruit List: ", len(fruitList))

	//declearing and initializing the array at the same time

	var vegList = [3]string{"Potato", "Tomato", "beans"}

	fmt.Println("VegList: ", vegList)

	//  IMP:    lenght stays the same as initialized size...no matter the no on elements inside an array
}
