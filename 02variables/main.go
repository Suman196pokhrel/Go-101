package main

import "fmt"

// beginning of letter is capital = > becomes public variable
const LOGINTOKES string = "daslfasdlfjsdlfjalsfdj"

func main() {
	var username string = "suman"
	fmt.Println(username)
	fmt.Printf("Variable of of type : %T \n", username)
	println("")

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable of of type : %T \n", isLoggedIn)
	println("")

	var smallValue uint64 = 98765321
	fmt.Println(smallValue)
	fmt.Printf("Variable is of Type : %T \n", smallValue)
	println("")

	var smallFloat float64 = 321.3232154625
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of Type : %T \n", smallFloat)
	println("")

	// default values and some aliases
	var anotherVariable1 string
	fmt.Println(anotherVariable1)
	println("")

	// implict type
	var webSite = "Test Your code"
	println(webSite)
	fmt.Printf("Variable is of type : %T \n", webSite)
	println("")

	// no var style
	numberOfUser := 1200
	fmt.Println(numberOfUser)
	println("")

	println(LOGINTOKES)
}
