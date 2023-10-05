package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Checking out the input feature in Go"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our SAAS platform")

	// comma ok Syntax (Error ok Syntax)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T, ", input)

}
