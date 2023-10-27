package main

import "fmt"



type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	
}

// (u User) is the receiver of the function , it associates the function with a specific type, allowing you to perform operations on instance of its type 
func (u User) GetStatus(){
	fmt.Println("Users status => ", u.Status)
}

func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is ", u.Email)
}



func main() {
	fmt.Println("Welcome to a session on methods")

	suman := User{"Suman Pokhrel", "emusk196@gmail.com", false, 24}
	
	suman.GetStatus()
	suman.NewMail()

	// Should be handeled by pointer later on  
	fmt.Printf("Updated User is => %+v",suman)
	
}