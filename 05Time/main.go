package main

import (
	"fmt"
	"time"
)

func main() {
	welcome := "Welcome to Go's default time moduel usage"
	fmt.Println(welcome)

	presentTime := time.Now()
	fmt.Println("Current time => ", presentTime)
	fmt.Println("Formatted Current time => ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	// Creating a time
	createdDate := time.Date(2020, time.July, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println("Created Date => ", createdDate)

}
