package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza App")
	fmt.Println("Please rate our app from 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	//numRating := input +1      this will through error beacuse we cant add int to string
	//hence we need string to int or float conversion

	// numRating, err := strconv.ParseFloat(input, 64) //second arg is the bit size of float

	// if err != nil {
	// 	fmt.Println("Error is: ", err)
	// } else {
	// 	fmt.Println("Numrating is: ", numRating+1)
	// }

	//THIS PIECE OF CODE GIVES ERROR :  strconv.ParseFloat: parsing "4\r\n": invalid syntax
	//the enter is also taken as input so we need to trim it using strings package

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //second arg is the bit size of float

	if err != nil {
		fmt.Println("Error is: ", err)
	} else {
		fmt.Println("Numrating is: ", numRating+1)
	}

}
