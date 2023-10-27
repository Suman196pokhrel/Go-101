package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to a session on SwitchCase")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Printf("Value of Dice is %d \n", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You Got One")

	case 2:
		fmt.Println("You Got Two")

	default:
		fmt.Println("Yout got something other than 1 and 2")

	}

}
