package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	//unlike array, we dont need to define size of slice

	var fruitList = []string{"Tomato", "Apple", "Mango"}

	fmt.Println("FruitList is: ", fruitList)

	//this is how we add members to slice, append takes slice in which we need to add members and returns the updated slice, then we assign the updated slice to our variable
	fruitList = append(fruitList, "Peach", "Grape")

	//append also gives us sliced slice--like in python

	fruitList = append(fruitList[1:3])

	//make keyword allows specific memory allocation

	highScores := make([]int, 4) //first arg -- type of var, second arg - size

	highScores[0] = 125
	highScores[1] = 554
	highScores[2] = 345
	highScores[3] = 894
	//highScores[4]=894   this will not work and will give out of bound errors since default allocation is 4 items only

	//with append method, memory is reallocated and hence we can fit more items

	highScores = append(highScores, 458, 945, 336)

	fmt.Println(highScores)

	//slice allow us to sort using sort package

	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores)) //provides boolean value

	//removing value from a slice when index is given

	courses := []string{"reactjs", "javascript", "swift", "python", "ruby"}

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...) //given remove item at index 2

	fmt.Println(courses)

}
