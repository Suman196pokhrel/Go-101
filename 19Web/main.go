package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev/"

func main() {
	fmt.Println("Welcome to a session on web requests")

	response, err := http.Get(url)

	if err != nil{
		panic(err)
	}

	fmt.Printf("RESPONSE of type %T  and data is =>\n", response)
	fmt.Println("RESPONSE => ", response)

	defer response.Body.Close()

	dataBytes, err :=io.ReadAll(response.Body)

	if err != nil{
		panic(err)
	}
	fmt.Println("Data is => ", string(dataBytes))
}