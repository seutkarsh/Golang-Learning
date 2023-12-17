package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to GO Lang Switch Case")

	//generating random number using math.rand

	// rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1 // gives no between 0-5 thats why add 1

	fmt.Println("Value of dice:", diceNumber)

	//switch case

	switch diceNumber {
	case 1:
		fmt.Println("Value of dice is 1 and you can open")
	case 2:
		fmt.Println("Value of dice is 2 and you can move 2 spots")
	case 3:
		fmt.Println("Value of dice is 3 and you can move 3 spots")
	case 4:
		fmt.Println("Value of dice is 4 and you can move 4 spots")

		//if we dont want to break and let it go through other cases too...use fallthrough
		fallthrough
		//now after hitting 4 it will also go for next case
	case 5:
		fmt.Println("Value of dice is 5 and you can move 5 spots")
	case 6:
		fmt.Println("Value of dice is 6 and you can move 6 spots and roll dice again")
	default:
		fmt.Println("What was that!!")
	}

}
