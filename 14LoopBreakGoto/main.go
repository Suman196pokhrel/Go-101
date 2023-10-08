package main

import "fmt"

func main() {
	fmt.Println("Welcome to sesstion on Loops in Golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println("Days => ", days)

	// LOOPS

	// TYPE-1
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// TYPE-2
	println("TYPE-2 loop")
	for i := range days {
		// range returns indexes not values
		fmt.Println(days[i])
	}

	// TYPE-2
	fmt.Println("TYPE-3")
	for index, day := range days {
		fmt.Printf("Index is %d, value is %v \n", index, day)
	}

	// BREAK- CONTINUE
	rougeValue := 1
	for rougeValue < 10 {
		if rougeValue == 5 {
			rougeValue++
			continue
		} else if rougeValue == 7 {
			goto sapiens
		} else if rougeValue == 9 {
			break
		}
		fmt.Println("Value is :", rougeValue)
		rougeValue++
	}

sapiens:
	fmt.Println("Read the book sapiens , its cool")

}
