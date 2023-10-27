package main

import "fmt"

func main() {
	fmt.Println("Welcome to a session on Pointer")

	myNumber := 25
	var pointerOne *int = &myNumber
	var pointerTwo = &myNumber

	fmt.Println("Value of empty pointer => ", pointerOne)  // direct refrence to memory address of myNumber
	fmt.Println("Value of empty pointer => ", &pointerOne) // address of pointerOne
	fmt.Println("Value of empty pointer => ", *pointerTwo) // value at provided memory location

	*pointerOne = *pointerOne * 2
	fmt.Println("New value is :", myNumber)

}
