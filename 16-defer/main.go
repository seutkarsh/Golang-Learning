package main

import "fmt"

func main() {

	//The defer keyword means the line will be stored in a stack (FIFO) and will be executed at the end of scope...in this case at the end of the function
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("hello")
	doDefer()
}

//The defer stack for main fn  -  World, One, Two     ----- Printing sequence - hello, two, one, world

//defer stack for dodefer fn - 0,1,2,3,4   -- printing for main - hello, 4,3,2,1,0, two,one,world

func doDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
