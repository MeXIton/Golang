package main

import "fmt"

func Hello() {
	fmt.Println("Hello Tonnam")
}

func plus(value1 int , value2 int) int{
	return value1+value2
}

func plus3value(value1 , value2 , value3 int ) int {
	return value1 + value2 + value3
}

func plus4value(price1 , price2 , price3 int) int {
	return price1 * price2 * price3
}

func main() {
	Hello()
	result := plus(1,2)
	fmt.Println("result = ",result)

	result2 := plus3value(5, 5 ,5)
	fmt.Println("result2 = ",result2)

	result3 := plus4value(10 , 5 , 10)
	fmt.Println("relust3 = ",result3)
}