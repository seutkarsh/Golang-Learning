package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const getUrl = "http://localhost:8000/get"
const jsonurl = "http://localhost:8000/post"
const formUrl = "http://localhost:8000/postform"

func main() {
	fmt.Println("Go Lang Requests")
	//performGetRequest(getUrl)
	// performPostRequestWithJsonData(jsonurl)
	performPostRequestWithForm(formUrl)

}

func performGetRequest(url string) {
	response, err := http.Get(url)
	checkNilError(err)

	defer response.Body.Close()

	fmt.Println("Response Status Code:", response.StatusCode)
	fmt.Println("Response Status Code:", response.ContentLength)

	data, err := io.ReadAll(response.Body)

	checkNilError(err)

	//another way of converting byte data into string

	var responseString strings.Builder

	byteCount, _ := responseString.Write(data)
	fmt.Println("Byte Count:", byteCount)
	fmt.Println(responseString.String())

}

func performPostRequestWithJsonData(url string) {
	//json data hack  -- new reader creates any data format as string
	requestBody := strings.NewReader(`
	{
		"name":"Utkarsh",
		"Job":"Developer"
	}
	`)

	response, err := http.Post(url, "application/json", requestBody)
	checkNilError(err)
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)

	checkNilError(err)

	var responseString strings.Builder

	byteCount, _ := responseString.Write(data)
	fmt.Println("Byte Count:", byteCount)
	fmt.Println(responseString.String())

}

func performPostRequestWithForm(myurl string) {
	//form data

	formData := url.Values{}

	formData.Add("name", "Utkarsh")
	formData.Add("age", "24")

	response, err := http.PostForm(myurl, formData)
	checkNilError(err)
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)

	checkNilError(err)

	var responseString strings.Builder

	byteCount, _ := responseString.Write(data)
	fmt.Println("Byte Count:", byteCount)
	fmt.Println(responseString.String())
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
