package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	// STRUCTS ARE USEAD AS AN ALTERNATIVES TO CLASSES IN GO
	// NO INHERITANCE IN GO
	fmt.Println("Welcome to a session of Structs")

	suman := User{"Suman Pokhrel", "emusk196@gmail.com", false, 24}
	fmt.Printf("USING STRUCT TO CREATE USER , %+v\n", suman)
	fmt.Println("Single struct value , ", suman.Email)

}
