package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/suman/25golangmongodb/router"
)

func main() {
	fmt.Println("MongoDb API Server")
	fmt.Println("Server booting up . . .")

	r := router.Router()

	log.Fatal(http.ListenAndServe(":8888", r))
	fmt.Println("Listening at port 8888")
}
