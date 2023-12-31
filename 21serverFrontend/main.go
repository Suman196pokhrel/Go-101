package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to a session on Interacting with web servers")
	const myGETUrl string = "http://localhost:8000/get"
	const myPOSTUrl string = "http://localhost:8000/post"
	const myPOSTFormUrl string = "http://localhost:8000/post"

	PerformGetRequest(myGETUrl)
	PerformPostJSONRequest(myPOSTUrl)
	PerformPostFormRequest(myPOSTFormUrl)

}

func PerformGetRequest(url string) {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code : ", response.StatusCode)
	fmt.Println("Content length is : ", response.ContentLength)

	// Creating My Own String Builder
	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count is : ", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJSONRequest(url string) {
	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"Golang Complete Tutorial",
			"price":110,
			"platform":"pyroprep"
		}
	`)

	respose, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer respose.Body.Close()
	content, _ := io.ReadAll(respose.Body)

	fmt.Println("POST RESPONSE => ", string(content))
}

func PerformPostFormRequest(myUrl string) {
	// fake form data
	data := url.Values{}
	data.Add("username", "Suman Pokhrel")
	data.Add("password", "testPassword")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println("POST FORM RESPONSE => ", string(content))

}
