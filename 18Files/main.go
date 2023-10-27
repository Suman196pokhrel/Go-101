package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to a session on Files in golang");
	context := "This needs to go into a file- Suman Pokhrel"

	file,err := os.Create("./myGoLangFile.txt")
	
	if err != nil{
		panic(err)
	}

	length, err := io.WriteString(file, context)
	if err!= nil{
		panic(err)
	}

	fmt.Println("length is :=> ", length)
	defer file.Close()

	readFile("./myGoLangFile.txt")

}


func readFile(fileName string){
	dataByte , err := ioutil.ReadFile(fileName)

	if err!=nil{
		panic(err)
	}

	fmt.Println("Text data inside the file is \n", string(dataByte))


}