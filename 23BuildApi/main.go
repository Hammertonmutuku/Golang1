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

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

//middleware, helper - file

func (c *Course) IsEmpty() bool {
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}
func main() {
	fmt.Println("Api")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}})
    //routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getoneCourse).Methods("GET")
	r.HandleFunc("/add-course", createOneCourse).Methods("POST")
    r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//contorollers - file

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Welcome to Api by learnCodeOnline</h1>"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getoneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("content-type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through courses , find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("content-type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")

	}

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}

	// generate unique id, string
	//append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("content-type", "application/json")

	// first - grab id from request
	params := mux.Vars(r)

	//loop, id, remove, add  with my ID

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
	//TODO: send a response when id is not found
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one course")
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	// loop id, remove(index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}
