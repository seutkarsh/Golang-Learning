package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time")

	presentTime := time.Now()

	fmt.Println(presentTime)

	//this results in this kind of format : 2023-12-12 13:59:00.0661157 +0530 IST m=+0.002860201

	//the format function takes special and specific values as parameters for formatting
	//EXACT THIS PARAMATERS
	// for date - 01-02-2006
	// for day - Monday
	// for time - 15:04:05
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	//creating date

	createdDate := time.Date(2000, time.August, 17, 23, 23, 0, 0, time.Local)

	fmt.Println("Created Date: ", createdDate)

}
