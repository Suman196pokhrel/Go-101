package main

import "fmt"

func main() {

	defer fmt.Println("Testing defer execution 1")
	defer fmt.Println("Testing defer execution 2")
	fmt.Println("Welcome to a session on defer")

	myDefer()

}



func myDefer(){
	for i:=0;i<5;i++{
		defer fmt.Println(i)
	}
}