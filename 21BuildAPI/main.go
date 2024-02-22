package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId   string  `json:"courseid"`
	CourseName string  `json:"coursename"`
	CourseType string  `json:"coursetype"`
	Price      int     `json:"price"`
	Author     *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleWares
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	e := mux.NewRouter()
	e.HandleFunc("/", serveHome)
	e.HandleFunc("/courses", getAllCourses).Methods("GET")
	e.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	e.HandleFunc("/course/delete/{id}", DeleteCourse).Methods("GET")
	e.HandleFunc("/delete", DeleteAllCourse).Methods("GET")
	e.HandleFunc("/course/update/{id}", updateCourse).Methods("PUT")
	e.HandleFunc("/course", addCourse).Methods("POST")
	http.ListenAndServe(":3000", e)
}

// Controllers - File

// serve home Route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearnCode Academy</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(courses) == 0 {
		json.NewEncoder(w).Encode("No courses found")
		return
	}
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// grab id
	params := mux.Vars(r)

	// loop though courses
	// find mathing id and return reponse
	for _, item := range courses {
		if item.CourseId == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given ID")
}

func addCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adding one course")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send a request body with the course details")
		return
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send a request body with the course details")
	}

	//  generate a unique ID and Convert it to string
	// append new course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range courses {
		if item.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given ID")
}

func DeleteAllCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	courses = []Course{}
	json.NewEncoder(w).Encode("Deleted all courses")
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range courses {
		if item.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course with ID: " + params["id"] + " has been deleted")
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given ID")
}
