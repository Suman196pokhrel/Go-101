package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to a session on Json Data in Go")
	EncodeJson()

	DecodeJson()

}

func EncodeJson() {
	pyroPrepCourses := []course{
		{"ReactRJ Bootcamp", 299, "PyroPrep", "testpassword", []string{"webdev", "js"}},
		{"VueJS Bootcamp", 199, "PyroPrepAlpha", "password", []string{"webdev", "js"}},
		{"AngularJS Bootcamp", 99, "PyroPrepBeta", "password123", nil},
	}

	// Package this data as JSON data
	finalJson, err := json.MarshalIndent(pyroPrepCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		 {
                "coursename": "ReactRJ Bootcamp",
                "Price": 299,
                "website": "PyroPrep",
                "tags": [
                        "webdev",
                        "js"
                ]
        }
	`)

	var pyroPrepCourse course

	checkVaile := json.Valid(jsonDataFromWeb)

	if checkVaile {
		fmt.Printf("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &pyroPrepCourse)
		fmt.Printf("%#v\n", pyroPrepCourse)
	} else {
		fmt.Printf("JSON WAS NOT VALID")
	}
}
