package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	//for loop

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// for looping through entire array or slice----- it gives index in variable

	for i := range days {
		fmt.Println(days[i])
	}

	//another way with value like for each   -- we can also replace _ for value we dont need

	for index, day := range days {
		fmt.Printf("Index Value:%v and Day Value: %v\n", index, day)
	}

	//while loop alike

	rougeValue := 1

	//if we want to execute something when rougeValue hit somewhere we use goto()

	for rougeValue < 10 {

		if rougeValue == 2 {
			goto lco
		}

		if rougeValue == 5 {
			//will break the for loop when rougeValue hit 5
			break
		}

		fmt.Println("Value is:", rougeValue)
		rougeValue++
	}

	//this is a label and it follows the statemenet....varname with colon is the syntax for label
lco:
	fmt.Println("Jumping to this")

}
