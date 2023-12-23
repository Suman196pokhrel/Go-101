// Write a program that takes a user’s name and age as input and displays a message on the console saying “Hello, [name]! You are [age] years old.”
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// var userName string
	// var userAge int

	// 1- USING SCAN  METHOD
	// fmt.Println("Enter your name : ")
	// fmt.Scan(&userName)
	// fmt.Println("Enter your age : ")
	// fmt.Scan(&userAge)

	// 2- USING SCANF
	// fmt.Println("Enter your name & Age : ")
	// fmt.Scanf("%d %s", &userName, &userAge)

	// 3- USING BUFFIO - it reads and writes data in chunks rather than in one byte at a time.
	// - This buffering can improve performance when dealing with large files or network streams
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter you name : ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading name : ", err)
		return
	}

	fmt.Println("Ënter your age : ")
	ageStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ërror reading age :", err)
		return
	}

	// Removing trailing newline & return
	ageStr = ageStr[:len(ageStr)-1]
	ageStr = ageStr[:len(ageStr)-1]

	// fmt.Println("AGE before parsing , ", ageStr)
	age, err := strconv.Atoi(ageStr)
	// fmt.Println("AGE after parsing , ", age)

	if err != nil {
		fmt.Print("Error while parsing the age str to int", err)
		return
	}

	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)

}
