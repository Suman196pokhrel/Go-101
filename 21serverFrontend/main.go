package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to a session on Interacting with web servers")
	const myUrl string = "http://localhost:8000/get"

	PerformGetRequest(myUrl)

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
