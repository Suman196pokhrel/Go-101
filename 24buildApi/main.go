package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for a course
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice string  `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// Mock DB
var courses []Course

// middlewares/ helpers
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""

}

func main() {
	fmt.Println("Welcome to the first api based proejct using go")

	router := mux.NewRouter()

	// endpoints
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/all", getAllCourses).Methods("GET")
	router.HandleFunc("/single/{courseId}", getSingleCourse).Methods("GET")
	router.HandleFunc("/createone", createOneCourse).Methods("POST")
	router.HandleFunc("/updateOnce/{courseId}", updateOneCourse).Methods("PUT")
	router.HandleFunc("/deleteOne/{courseId}", deleteOneCourse).Methods("DELETE")

	// server
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Controllers - file

// serve Home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome Go DigicalEd</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getSingleCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get single course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range courses {
		if item.CourseId == params["courseId"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with given id ")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create single course")
	w.Header().Set("Content-Type", "application/json")

	// if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// if body is {}
	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No Data provided")
		return
	}

	// generate uniqueId , convert it into string
	// append new course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// Grab Id
	params := mux.Vars(r)
	var course Course

	for index, item := range courses {
		if item.CourseId == params["courseId"] {
			// REmove
			courses = append(courses[:index], courses[index+1:]...)

			// Add new
			json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["courseId"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
		}
	}

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	params := mux.Vars(r)

	for index, item := range courses {
		if params["courseId"] == item.CourseId {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Deleted course with given Id")
			return
		}
	}

	json.NewEncoder(w).Encode("Could not fuind course with given id")
}
