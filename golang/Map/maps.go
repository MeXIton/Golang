package main

import "fmt"

var product = make(map[string]float64)

func main() {
	 //Add 
	 product["Macbook"] = 40000
	 product["Iphone"] = 40000
	 fmt.Println("product = ",product)

	 //Delete
	 delete(product , "Iphone")
	 fmt.Println("product = ",product)

	 //Update 
	 product["Macbook"] = 400007
	 fmt.Println("product = ",product)

	 value1 := product["Iphone"]
	 fmt.Println("value = ", value1)

	 courseName := map[string]string{"101":"Lava","102": "Python"}
	 fmt.Println("courseName = ", courseName)

}