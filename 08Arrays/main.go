package main

import "fmt"

func main() {
	fmt.Println("Welcome to a Session on Arrays")

	// INITIALIZE WITHOUT VALUES
	var books [3]string
	books[0] = "The Psychology of money"
	books[2] = "Sapiens"

	fmt.Println("The Booklist is =>", books)
	fmt.Println("The books length =>", len(books))

	// INITIALIZE WITH VALUES
	vegList := [3]string{"potato", "beans", "mushroom"}
	fmt.Println("The vegetable list is => ", vegList)

}
