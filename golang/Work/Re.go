package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

func coursehandler(w http.ResponseWriter, r *http.Request) {
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
	http.HandleFunc("/course", coursehandler)
	http.ListenAndServe(":5000", nil)
}