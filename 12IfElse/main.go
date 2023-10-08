package main

import "fmt"

func main() {
	fmt.Println("Welcome to a session on If/Else")

	loginCount := 10

	var result string

	// REGULAR
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Power User"
	} else {
		result = "Mid User"
	}

	fmt.Printf("Current user is a %v \n", result)

	// COMPUTATION ON CONDITION
	if 9%2 == 0 {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("Number is Odd")
	}

	// INITIALIZE AND CHECK VALUE ON THE GO
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

}
