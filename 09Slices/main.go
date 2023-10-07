package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to session on Slices")

	var books = []string{"RDPD", "Sapiens", "Man's search for meaning"}
	fmt.Printf("Type of Books is %T \n", books)

	books = append(books, "Driving License")
	fmt.Println("Updated books Slice =>> ", books)

	books = append(books[1:2])
	fmt.Println("Sliced books Slice =>> ", books)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 944
	highScores[2] = 254
	highScores[3] = 264

	highScores = append(highScores, 555, 444, 333, 222)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println("Sorted Highscores => ", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove value from Slices based on index
	var courses = []string{"reactjs", "JAVAscript", "Kotlin", "Python", "Ruby"}
	var index int = 2
	fmt.Println("Courses => ", courses)
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Updated coureses= >>> ", courses)

}
