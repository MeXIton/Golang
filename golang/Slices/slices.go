package main

import "fmt"

var courseName []string

func main() {
	courseName = []string{"Lava", "Python"}
	courseName = append(courseName, "Golang", "Javascript" , "HTML" , "CSS")
	fmt.Println(courseName)

	courseWeb := courseName[4:4]
	fmt.Println(courseWeb)
}
