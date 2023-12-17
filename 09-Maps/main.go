package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps")

	// map[key_type]value_type
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages)
	fmt.Println(languages["JS"])

	//to delete a value with keyname

	delete(languages, "RB")

	fmt.Println(languages)

	//iterating over a map

	for key, value := range languages {
		fmt.Println(key, value)
	}
}
