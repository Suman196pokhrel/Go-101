package main

import "fmt"

func main() {
	fmt.Println("Welcome to a session on Maps in Go")

	// HASH TABLE / HASH MAP
	languages := make(map[string]string)

	languages["JS"] = "JAVASCRIPT"
	languages["PY"] = "PYTHON"
	languages["C"] = "C"
	languages["K"] = "KOTLIN"

	fmt.Println("Languages => ", languages)
	fmt.Println("JS => ", languages["JS"])

	delete(languages, "K")
	fmt.Println("Languages => ", languages)

}
