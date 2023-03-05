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

// models for course

type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"-"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json: "fullname"`
	Website  string `json: "website"`
}

// fake db

var courses []Course

// middleware, or  helper-> file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""

}
func main() {
	fmt.Println("Build API")
	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "2", CourseName: "GO Backend", CoursePrice: 230, Author: &Author{FullName: "Bikash", Website: "YT"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "Full Stack Developer", CoursePrice: 231, Author: &Author{FullName: "Bikash", Website: "YT"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourse).Methods("GET")
	r.HandleFunc("/course", addCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	// listen
	log.Fatal(http.ListenAndServe(":4000", r))

}

// controller file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Home Page for Routers</h1>"))

}

func getCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting List of all Courses.")
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("List of all courses: ", courses)
	json.NewEncoder(w).Encode(courses)

}

func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Course by course_id: ", r.FormValue("course_id"))
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)
	// loop through course, find matching course and get detail of course

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found for this course id.")
	return
}

func addCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create course")
	w.Header().Set("Content-Type", "application/json")

	// TODO: check duplicate course name
	// : loop, check course name match or not, json response if match
	// --- do here ---

	// what if body is empty?
	if r.Body != nil {
		json.NewEncoder(w).Encode("Please fill the body with the course")
	}
	// what if data is like-{}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No  data found.")

	}

	// generate unique id
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	// course.CourseId = rand.Intn(100)
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create course")
	w.Header().Set("Content-Type", "application/json")

	//  grab id from request body
	params := mux.Vars(r)

	// loop to body data, get id, remove by slice and add again by change data

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return

		}

	}
	// TODO: send a response if id not found

}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete course by id")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// loop, id, remove(index, index)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break

		}
	}
	json.NewEncoder(w).Encode("course deleted successfully.")
	return
}
