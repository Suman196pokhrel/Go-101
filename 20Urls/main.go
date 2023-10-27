package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=laskdjlaskjdl"

func main() {
	fmt.Println("Welcome to a session on handeling Urls")
	fmt.Println("Our Urls => ", myUrl)

	// parsing
	result, _ := url.Parse(myUrl)

	fmt.Println("Our result =>", result)
	fmt.Println("Scheme =>", result.Scheme)
	fmt.Println("HOST =>", result.Host)
	fmt.Println("Path =>", result.Path)
	fmt.Println("Port =>", result.Port())
	fmt.Println("Raw Query =>", result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Typ of Query params are : %T \n %v \n", qparams, qparams)

	for _, val := range qparams {
		fmt.Println(val)
	}

	// Constructing the URL
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "localhost",
		Path:     "/testNft",
		RawQuery: "user=suman",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("Creating a URl => ", anotherUrl)

}
