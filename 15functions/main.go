package main

import "fmt"

func main() {

	fmt.Println("Welcome to a session on functions");
	greeter();
	
	result := adder(3, 5)
	fmt.Println("Result is : ", result)

	proResult,myMessage := proAdder(2,5,6,7,8)
	fmt.Println("Pro result is  : ", proResult)
	fmt.Println("Pro message is  : ", myMessage)

}


func proAdder (values ...int) (int,string){
	total := 0
	for _, value := range values{
		total += value
	}

	return total ,"Pro result incoming"
}


func adder (valOne int, valTwo int) int{
	return (valOne + valTwo)
}

func greeter()  {
	fmt.Println("Hello from golang")
}