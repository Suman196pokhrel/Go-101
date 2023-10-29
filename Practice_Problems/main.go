package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to go basics")

	// displayNameUsingReader()

	// displayNameUsingScan()

	// showAddress()

	// multiLineComments()

	// takeMultipleInputs()

	// findAge()

	// printNumbers()

	// createArray()

	// createArrayWithValues()

	// creatingSlieces()

	// slicesFunctions()

	loopsUsingRange()
}

func displayNameUsingReader() {
	fmt.Printf("Running Program to Display User name \n")

	// Creating a reader that reads from standard Input
	reader := bufio.NewReader(os.Stdin)

	// Read a line of text(until the Enter Key is Pressed)
	input, _ := reader.ReadString('\n')

	// Removing any leading or Traling Spaces
	input = strings.TrimSpace(input)

	fmt.Printf("You Enteres : %s\n", input)
}

func displayNameUsingScan() {
	var userNameLine string
	var userNameFormatted string

	// Reads line of text until newline char
	fmt.Println("Enter your Username => ")
	fmt.Scanln(&userNameLine)

	// reads input based on formatted strings
	fmt.Printf("Enter you Username for formatted input => \n")
	fmt.Scanf("%s", &userNameFormatted)

	fmt.Printf("You Entered %s for normal input \n", userNameLine)
	fmt.Printf("You Entered %s for formatted input\n", userNameFormatted)
}

func showAddress() {
	var country, city, street string

	fmt.Println("Enter your country name")
	fmt.Scanf("%s", &country)

	fmt.Println("Enter your City name")
	fmt.Scanf("%s", &city)

	fmt.Println("Enter your street name")
	fmt.Scanf("%s", &street)

	address := street + "," + city + "," + country
	fmt.Println("you complete address is , ", address)
}

func multiLineComments() {
	// Single line comment

	/*

		Now this is a
		multiline comment

	*/
	test := "hello"
	fmt.Println(test)
}

func takeMultipleInputs() {

	// Only workd for single keyworded input
	var name, address string
	var age int
	var isAdult bool = false

	fmt.Println("Enter your full name")
	fmt.Scanf("%s", &name)

	fmt.Println("Enter your age")
	fmt.Scanf("%d", &age)

	fmt.Println("Enter your Address")
	fmt.Scanf("%s", &address)

	if age > 18 {
		isAdult = true
	}

	fmt.Printf("Here's Your Info:\n")
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Address: %s\n", address)
	fmt.Printf("Is Adult: %t\n", isAdult)
}

func findAge() {
	var year, month, day string

	fmt.Println("Enter the year you were born")
	fmt.Scanf("%s", &year)

	fmt.Println("Enter the month you were born")
	fmt.Scanf("%s", &month)

	fmt.Println("Enter the day you were born")
	fmt.Scanf("%s", &day)

	// converting to integers
	yearInt, _ := strconv.Atoi(year)
	monthInt, _ := strconv.Atoi(month)
	dayInt, _ := strconv.Atoi(day)

	fmt.Printf("Year: %d\n", yearInt)
	fmt.Printf("Month: %d\n", monthInt)
	fmt.Printf("Day: %d\n", dayInt)

	// todays time
	today := time.Now()

	// Creating a time object for birthdate
	birthDate := time.Date(yearInt, time.Month(monthInt), dayInt, 0, 0, 0, 0, time.UTC)
	age := today.Year() - birthDate.Year()

	fmt.Printf("Your age is : %d \n", age)
}

func printNumbers() {
	var n int

	fmt.Println("Enter the number of numbers you want")
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		fmt.Println(i)
	}

	fmt.Println("Task finished")
}

func createArray() {
	var books [3]string

	books[0] = "Psychology of money"
	books[1] = "Sapiens"
	books[2] = "Man's search for meaning of life"

	fmt.Println("Your books Array => ", books)
}

func createArrayWithValues() {
	books := [3]string{"psychology of money", "Sapiens", "Man's search for meaning"}

	fmt.Println("Your books array with explict values are => ", books)
}

func creatingSlieces() {
	books := []string{"Psychology of money", "sapiens", "Man's search for meaning"}

	fmt.Println("Books Slices => ", books)
}

func slicesFunctions() {
	var numOfUsers int
	var username string
	var books []string

	fmt.Print("Lets create users slices")
	fmt.Println()

	fmt.Println("Enter the number of Users you want")
	fmt.Scanf("%d", &numOfUsers)

	// appending elements to the books slices
	for i := 0; i < numOfUsers; i++ {
		fmt.Printf("Enter username for user %d : ", i)
		fmt.Scanf("%s", &username)
		books = append(books, username)
	}

	fmt.Println("Woudl you like to Create a subslice")

	fmt.Printf("Your User slice %v \n", books)

	fmt.Printf("SubSlice test %v", books[1:2])

}

func loopsUsingRange() {
	books := []string{"Psychology of money", "spaiens", "man's search for meaning"}

	for index, item := range books {
		fmt.Printf("Index is %d and item is %s \n", index, item)
	}
}
