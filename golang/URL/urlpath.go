package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Course struct {
	CourseId   int     `json:"id"`
	CourseName string  `json:"name"`
	Price      float64 `json:"price"`
	Instructor string  `json:"instructor"`
}

var CourseList []Course

func init() {
	CourseJSON := `[
	{
		"id": 1,
		"name": "Python",
		"price": 2590,
		"instructor": "Naruepai"
	},
	{
		"id": 2,
		"name": "Golang",
		"price": 2591,
		"instructor": "Namto1"
	},
	{
		"id": 3,
		"name": "Rust",
		"price": 2592,
		"instructor": "Tonnam2"
	}
]`
	err := json.Unmarshal([]byte(CourseJSON), &CourseList)
	if err != nil {
		log.Fatal(err)
	}
}

func getNextID() int {
	highestID := -1 // Correct the typo in "highestID"
	for _, course := range CourseList {
		if highestID < course.CourseId {
			highestID = course.CourseId
		}
	}
	return highestID + 1 // This should be outside the loop
}

func findID(ID int) (*Course, int) {
	for i, course := range CourseList {
		if course.CourseId == ID {
			return &course, i
		}
	}	
	return nil, 0
}

func coursehandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegmant := strings.Split(r.URL.Path, "course/")
	ID, err := strconv.Atoi(urlPathSegmant[len(urlPathSegmant)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	
	course, ListItemIndex:= findID(ID)
	if course == nil {
		http.Error(w, fmt.Sprintf("No Course with ID %d", ID),http.StatusNotFound)
	}
	switch r.Method {
	case http.MethodGet: 
	courseJson, err := json.Marshal(course)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		}
		w.Header().Set("Content-Type","application/json")
		w.Write(courseJson)

	case http.MethodPut:
		var updateCourse Course
		byteBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(byteBody, &updateCourse) 
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if updateCourse.CourseId != ID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		course = &updateCourse
		CourseList[ListItemIndex] = *course
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func courseshandler(w http.ResponseWriter, r *http.Request) {
	coursejson, err := json.Marshal(CourseList)
	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
}
		w.Header().Set("Content-Type", "application/json")
		w.Write(coursejson)
	case http.MethodPost:
		var newCourse Course
		Bodybyte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(Bodybyte, newCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		CourseList = append(CourseList, newCourse)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func main() {
	http.HandleFunc("/course/", coursehandler)
	http.HandleFunc("/course", courseshandler)
	http.ListenAndServe(":5000", nil)
}