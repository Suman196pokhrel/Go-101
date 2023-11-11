package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("A full features golang app integrated with mongoDb")

	r := mux.NewRouter()

	http.ListenAndServe(":8888", r)
}
