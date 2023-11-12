package main

import "fmt"

var product [4]string

func main() {
	product[0] = "Apple"
	product[1] = "Banana"
	product[2] = "iPhone"
	product[3] = "iPad"
	fmt.Println(product)

	price := [4]float64{4000 , 5000 , 6000 , 7000}
	fmt.Println(price)

}