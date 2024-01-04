package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Json")
	//EncodeJson()
	DecodeJson()
}

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              //this means not show in json conversion
	Tags     []string `json:"tags,omitempty"` //omitempty means dont show if empty
}

func EncodeJson() {

	courses := []course{
		{"React Js", 299, "Coursera", "abc@123", []string{"web-dev", "online"}},
		{"Vue Js", 159, "Udemy", "bvc@123", []string{"web-dev", "frontend"}},
		{"React Native", 299, "Coursera", "xyz@123", nil},
	}

	//package this data as JSON data

	//finalJson, err := json.Marshal(courses)
	// output  -- [{"Name":"React Js","Price":299,"Platform":"Coursera","Password":"abc@123","Tags":["web-dev","online"]},{"Name":"Vue Js","Price":159,"Platform":"Udemy","Password":"bvc@123","Tags":["web-dev","frontend"]},{"Name":"React Native","Price":299,"Platform":"Coursera","Password":"xyz@123","Tags":null}]

	//to format better..use marshalIndent

	finalJson, err := json.MarshalIndent(courses, "", "\t")
	//output -- without alias in struct
	//[
	//         {
	//                 "Name": "React Js",
	//                 "Price": 299,
	//                 "Platform": "Coursera",
	//                 "Password": "abc@123",
	//                 "Tags": [
	//                         "web-dev",
	//                         "online"
	//                 ]
	//         },
	//         {
	//                 "Name": "Vue Js",
	//                 "Price": 159,
	//                 "Platform": "Udemy",
	//                 "Password": "bvc@123",
	//                 "Tags": [
	//                         "web-dev",
	//                         "frontend"
	//                 ]
	//         },
	//         {
	//                 "Name": "React Native",
	//                 "Price": 299,
	//                 "Platform": "Coursera",
	//                 "Password": "xyz@123",
	//                 "Tags": null
	//         }
	// ]

	//output after alias in struct
	// 	[
	//         {
	//                 "coursename": "React Js",
	//                 "Price": 299,
	//                 "website": "Coursera",
	//                 "tags": [
	//                         "web-dev",
	//                         "online"
	//                 ]
	//         },
	//         {
	//                 "coursename": "Vue Js",
	//                 "Price": 159,
	//                 "website": "Udemy",
	//                 "tags": [
	//                         "web-dev",
	//                         "frontend"
	//                 ]
	//         },
	//         {
	//                 "coursename": "React Native",
	//                 "Price": 299,
	//                 "website": "Coursera"
	//         }
	// ]
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	//create a json from byte data
	jsonFromWeb := []byte(
		`{
                "coursename": "React Js",
                "Price": 299,
                "website": "Coursera",
                "tags": [
                        "web-dev",
                        "online"
                ]
        }`)

	// store json in a var of type course (Struct that we waht our json to have -- like an interface)
	var courses course

	//validate json
	jsonValidity := json.Valid(jsonFromWeb)

	if jsonValidity {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonFromWeb, &courses)
		fmt.Printf("%#v\n", courses)
	} else {
		fmt.Println("JSON IS NOT VALID")
	}

	// if we want json into key value pair instead of a struct

	var onlineCourses map[string]interface{}
	json.Unmarshal(jsonFromWeb, &onlineCourses)
	fmt.Printf("%#v\n", onlineCourses)

	for k, v := range onlineCourses {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}

}
